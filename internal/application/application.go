package model

import (
	"awesomeProject/internal/infrastructure/persistence"
	sys "awesomeProject/internal/proto/health"
	hrPb "awesomeProject/internal/proto/hr"
	resumePb "awesomeProject/internal/proto/resume"
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
	resumePb.UnimplementedResumeServiceServer
	hrPb.UnimplementedHrServiceServer
}

func NewApplication(config Config, logger *log.Logger, DB *sql.DB, S3 *persistence.S3) *Application {
	return &Application{
		config:      config,
		logger:      logger,
		persistence: persistence.NewRepositories(DB, S3),
	}
}
