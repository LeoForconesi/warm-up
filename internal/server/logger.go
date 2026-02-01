package server

import (
	"go.uber.org/zap"
)

func NewLogger(env string) *zap.Logger {
	if env == "production" {
		logger, _ := zap.NewProduction()
		return logger
	}
	logger, _ := zap.NewDevelopment()
	return logger
}
