package logging

import "go.uber.org/zap"

// Init Logger
func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	return cfg.Build()
}
