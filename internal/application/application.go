package model

import (
	sys "awesomeProject/internal/proto/health"
	studentPb "awesomeProject/internal/proto/student"
	"log"
)

type Config struct {
	Port    string
	Env     string
	Version string
	DB      struct {
		DSN         string
		MaxOpenConn int
		MaxIdleConn int
		MaxIdleTime string
	}
}

type Application struct {
	config Config
	logger *log.Logger
	sys.UnimplementedHealthCheckServer
	studentPb.UnimplementedStudentServiceServer
}

func NewApplication(config Config, logger *log.Logger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}
