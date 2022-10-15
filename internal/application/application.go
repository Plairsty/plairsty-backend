package model

import (
	"awesomeProject/internal/infrastructure/persistence"
	"awesomeProject/internal/infrastructure/service"
	jobApplicationPb "awesomeProject/internal/proto/application"
	authPb "awesomeProject/internal/proto/auth"
	certificatePb "awesomeProject/internal/proto/certificates"
	sys "awesomeProject/internal/proto/health"
	hrPb "awesomeProject/internal/proto/hr"
	internshipPb "awesomeProject/internal/proto/internship"
	projectPb "awesomeProject/internal/proto/project"
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
	config          Config
	logger          *log.Logger
	persistence     *persistence.Repositories
	jwtManager      *service.JwtManager
	accessibleRoles map[string][]string
	sys.UnimplementedHealthCheckServer
	authPb.UnimplementedAuthServiceServer
	studentPb.UnimplementedStudentServiceServer
	resumePb.UnimplementedResumeServiceServer
	hrPb.UnimplementedHrServiceServer
	jobApplicationPb.UnimplementedJobApplicationServiceServer
	certificatePb.UnimplementedCertificateServiceServer
	internshipPb.UnimplementedInternshipServiceServer
	projectPb.UnimplementedProjectServiceServer
}

func NewApplication(
	config Config,
	logger *log.Logger,
	DB *sql.DB,
	S3 *persistence.S3,
	jwtManager *service.JwtManager,
	accessibleRoles map[string][]string,
) *Application {
	return &Application{
		config:          config,
		logger:          logger,
		persistence:     persistence.NewRepositories(DB, S3),
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}
