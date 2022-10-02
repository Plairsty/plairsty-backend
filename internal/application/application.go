package model

import (
	sys "awesomeProject/internal/proto/health"
	"context"
	"log"
)

type Config struct {
	Port    string
	Env     string
	Version string
}

type Application struct {
	config Config
	logger *log.Logger
	sys.UnimplementedHealthCheckServer
}

func NewApplication(config Config, logger *log.Logger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}

func (app *Application) Check(
	_ context.Context,
	_ *sys.HealthCheckRequest,
) (*sys.HealthCheckResponse, error) {
	return &sys.HealthCheckResponse{
		Status:      "available",
		Environment: app.config.Env,
		Version:     app.config.Version,
	}, nil
}
