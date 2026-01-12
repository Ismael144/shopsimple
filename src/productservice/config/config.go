package config

import (
	"log"
	"os"
)

type Config struct {
	GRPCAddr    string
	DatabaseURL string
	JaegerURL   string
}

func LoadConfig() Config {
	return Config{
		GRPCAddr:    getEnv("GRPC_ADDR", ":50051"),
		DatabaseURL: mustEnv("DATABASE_URL"),
		JaegerURL:   mustEnv("JAEGER_URL"),
	}
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env var: %s", key)
	}
	return v
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
