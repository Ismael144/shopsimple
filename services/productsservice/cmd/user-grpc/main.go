package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ismael144/productsservice/config"
	"github.com/Ismael144/productsservice/internal/application"
	"github.com/Ismael144/productsservice/internal/infrastructure/db"
	"github.com/Ismael144/productsservice/internal/infrastructure/repository"
	grpcTransport "github.com/Ismael144/productsservice/internal/transport/grpc"
	"github.com/Ismael144/productsservice/internal/transport/grpc/interceptors"
)

func main() {
	cfg := config.LoadConfig()

	// Database
	gormDB, err := db.NewPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Infrastructure
	productsRepo := repository.NewProductsRepository(gormDB)

	// Application
	productsService := application.NewProductsService(
		productsRepo,
	)

	// Initialize server
	server, err := grpcTransport.NewServer(
		cfg.GRPCAddr,
		productsService,
		interceptors.RequestID(),
		interceptors.Logging(),
	)

	if err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}

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
