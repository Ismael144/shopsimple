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
	"github.com/Ismael144/productservice/migrations"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file...")
	}

	ctx := context.Background()
	cfg := config.LoadConfig()

	// Database
	client, err := db.NewMongoClient(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %+v", err)
	}

	defer func() {
		_ = client.Disconnect(ctx)
	}()

	// Mongo Db
	db := db.NewMongoDatabase(client, os.Getenv("MONGO_DB"))

	// Auto Migrate Models
	if cfg.RunMigrations == "true" {
		if err := migrations.EnsureProductIndexes(ctx, db); err != nil {
			log.Fatalf("Failed to run migrations: %+v", err)
		}
	}

	// Infrastructure
	productsRepo := repository.NewMongoProductRepository(db)
	logger, _ := logging.New()
	defer logger.Sync()

	shutdown, err := telemetry.InitTracer("product-service", cfg.JaegerURL)
	if err != nil {
		logger.Fatal("failed to init tracer", zap.Error(err))
	}
	defer shutdown(ctx)

	// Application
	productservice := application.NewProductService(
		productsRepo,
	)

	// Initialize server
	server, err := grpcTransport.NewServer(
		cfg.GRPCAddr,
		productservice,
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
