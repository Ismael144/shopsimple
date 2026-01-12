package config

import (
	"log"
	"os"
)

type Config struct {
	GRPCAddr          string
	RedisAddr         string
	JaegarURL         string
	ProductServerAddr string
}

func LoadConfig() Config {
	return Config{
		GRPCAddr:          getEnv("GRPC_ADDR", ":50052"),
		RedisAddr:         mustEnv("REDIS_ADDR"),
		JaegarURL:         mustEnv("JAEGER_URL"),
		ProductServerAddr: mustEnv("PRODUCTSERVICE_ADDR"),
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
