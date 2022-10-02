package model

import (
	sys "awesomeProject/internal/proto/health"
	"context"
)

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
