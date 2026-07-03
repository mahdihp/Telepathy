package logger

import (
	"go.uber.org/zap"
)

func NewDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewProduction() (*zap.Logger, error) {
	return zap.NewProduction()
}

func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.Encoding = "json"

	return cfg.Build()
}

type UserHandler struct {
	//service *UserService
	logger *zap.Logger
}
