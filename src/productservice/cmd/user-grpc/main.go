package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ismael144/productservice/config"
	"github.com/Ismael144/productservice/internal/application"
	"github.com/Ismael144/productservice/internal/infrastructure/db"
	"github.com/Ismael144/productservice/internal/infrastructure/logging"
	"github.com/Ismael144/productservice/internal/infrastructure/repository"
	"github.com/Ismael144/productservice/internal/infrastructure/telemetry"
	grpcTransport "github.com/Ismael144/productservice/internal/transport/grpc"
	"github.com/Ismael144/productservice/internal/transport/grpc/interceptors"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	// Database
	gormDB, err := db.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Infrastructure
	productsRepo := repository.NewProductsRepository(gormDB)
	categoriesRepo := repository.NewProductCategoryRepository(gormDB)
	ctx := context.Background()
	logger, _ := logging.New()
	defer logger.Sync()

	shutdown, err := telemetry.InitTracer("product-service")
	if err != nil {
		logger.Fatal("failed to init tracer", zap.Error(err))
	}
	defer shutdown(ctx)

	// Application
	productservice := application.NewProductservice(
		productsRepo,
	)

	categoryservice := application.NewProductCategoryService(
		categoriesRepo,
	)

	// Initialize server
	server, err := grpcTransport.NewServer(
		cfg.GRPCAddr,
		productservice,
		categoryservice,
		interceptors.RequestIDInterceptor(),
		interceptors.LoggingInterceptor(logger),
	)

	if err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}

	// Start prometheus metrics server
	_, metricsHandler, err := telemetry.InitMetrics()
	if err != nil {
		logger.Fatal("metrics init failed", zap.Error(err))
	}

	go func() {
		logger.Info("Metrics server started", zap.String("addr", ":9090"))
		http.ListenAndServe(":9090", metricsHandler)
	}()

	// Grpc server
	go func() {
		log.Printf("starting gRPC server on %s", cfg.GRPCAddr)
		if err := server.Start(); err != nil {
			log.Fatalf("GRPC Server error: %v", err)
		}
	}()

	// Graceful shutdown
	waitForShutdown(server)
}

func waitForShutdown(server *grpcTransport.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	sig := <-ch
	log.Printf("received signal %s, shutting down", sig)

	server.Stop()
	log.Printf("Server stopped gracefully...")
}
