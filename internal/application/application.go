package model

import (
	"awesomeProject/internal/infrastructure/persistence"
	sys "awesomeProject/internal/proto/health"
	studentPb "awesomeProject/internal/proto/student"
	"database/sql"
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
	config      Config
	logger      *log.Logger
	persistence *persistence.Repositories
	sys.UnimplementedHealthCheckServer
	studentPb.UnimplementedStudentServiceServer
}

func NewApplication(config Config, logger *log.Logger, DB *sql.DB) *Application {
	return &Application{
		config:      config,
		logger:      logger,
		persistence: persistence.NewRepositories(DB),
	}
}
