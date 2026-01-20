package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ismael144/cartservice/config"
	"github.com/Ismael144/cartservice/internal/application"
	"github.com/Ismael144/cartservice/internal/infrastructure/db"
	"github.com/Ismael144/cartservice/internal/infrastructure/logging"
	"github.com/Ismael144/cartservice/internal/infrastructure/repository"
	"github.com/Ismael144/cartservice/internal/infrastructure/telemetry"
	grpcTransport "github.com/Ismael144/cartservice/internal/transport/grpc"
	"github.com/Ismael144/cartservice/internal/transport/grpc/clients"
	"github.com/Ismael144/cartservice/internal/transport/grpc/interceptors"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file...")
	}

	cfg := config.LoadConfig()
	ctx := context.Background()

	// Initialize db
	rdb := db.NewRedis(cfg.RedisAddr)

	// Infrastructure
	cartRepo := repository.NewCartRepositoryRedis(rdb)
	logger, _ := logging.New()
	defer logger.Sync()

	shutdown, err := telemetry.InitTracer("cart-service", cfg.JaegarURL)
	if err != nil {
		logger.Fatal("failed to init tracer", zap.Error(err))
	}
	defer shutdown(ctx)

	// Application
	productClient, err := clients.NewProductServiceServerClient(cfg.ProductServerAddr)
	if err != nil {
		log.Fatalf("Failed to initialize products service client")
	}
	currencyClient, err := clients.NewCurrencyServiceServerClient(cfg.CurrencyServerAddr)

	cartservice := application.NewCartService(
		cartRepo,
	)

	server, err := grpcTransport.NewServer(
		cfg.GRPCAddr,
		cartservice,
		productClient,
		currencyClient,
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
	log.Printf("Server stop gracefully...")
}
