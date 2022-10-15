package main

import (
	"awesomeProject/internal/application"
	"awesomeProject/internal/infrastructure/persistence"
	"awesomeProject/internal/infrastructure/service"
	jobApplicationPb "awesomeProject/internal/proto/application"
	authPb "awesomeProject/internal/proto/auth"
	certificatePb "awesomeProject/internal/proto/certificates"
	sys "awesomeProject/internal/proto/health"
	hrPb "awesomeProject/internal/proto/hr"
	resumePb "awesomeProject/internal/proto/resume"
	studentPb "awesomeProject/internal/proto/student"
	"awesomeProject/utils"
	"context"
	"database/sql"
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "1.0.0"

func init() {
	utils.DevBanner()
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	var cfg model.Config
	cfg.Version = version
	flag.StringVar(&cfg.Port, "port", ":4000", "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("PLAIRSTY_DB_DSN"), "PostgreSQL DSN")

	// Read the connection pool settings from command-line flags into the config struct.
	// Notice the default values that we're using?
	flag.IntVar(&cfg.DB.MaxOpenConn, "db-max-open-conn", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConn, "db-max-idle-conn", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Open database connection
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(db)
	// S3 config
	S3 := &persistence.S3{
		Region:   os.Getenv("REGION"),
		Bucket:   os.Getenv("S3_BUCKET"),
		Key:      "folder1/",
		ApiId:    os.Getenv("AWS_ACCESS_KEY_ID"),
		ApiToken: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	log.Println("API ID: ", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	jwtManager := service.NewJwtManager(
		os.Getenv("PLAIRSTY_JWT_SECRET"),
		os.Getenv("PLAIRSTY_JWT_ISSUER"),
		time.Hour*1,
	)

	// After all configuration is done, we can create a new application instance.
	app := model.NewApplication(
		cfg,
		logger,
		db,
		S3,
		jwtManager,
		accessibleRoles(),
	)

	// Register gRPC server
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		logger.Fatal(err)
	}
	log.Println("Starting server on port", cfg.Port)
	defer func(conn net.Listener) {
		log.Println("Closing server on port", cfg.Port)
		err := conn.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(lis)

	// Create a new gRPC server with the UnaryServerInterceptor and StreamServerInterceptor 🦹
	server := grpc.NewServer(
		grpc.UnaryInterceptor(app.Unary()),
		grpc.StreamInterceptor(app.Stream()),
	)

	// Register your services here 🥵
	// If you found an error stating that certain method is not implemented,
	// Check if you have implemented the method in your service implementation
	// And then check if you have added unimplemented method in your service struct
	sys.RegisterHealthCheckServer(server, app)
	authPb.RegisterAuthServiceServer(server, app)
	studentPb.RegisterStudentServiceServer(server, app)
	resumePb.RegisterResumeServiceServer(server, app)
	hrPb.RegisterHrServiceServer(server, app)
	jobApplicationPb.RegisterJobApplicationServiceServer(server, app)
	certificatePb.RegisterCertificateServiceServer(server, app)

	// register reflection service on gRPC server.
	reflection.Register(server)

	// Get service info
	//for _, service := range server.GetServiceInfo() {
	//	go func(s grpc.ServiceInfo) {
	//		for _, method := range s.Methods {
	//			log.Println("Invoking Method:", method.Name)
	//		}
	//	}(service)
	//}

	// background goroutine for graceful shutdown
	go func() {
		// Create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)
		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel. Any other signals will not be caught by
		// signal.Notify() and will retain their default behavior.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// Read the signal from the quit channel. This code will block until a signal is
		// received.
		s := <-quit
		log.Println("Shutting down server...", map[string]interface{}{"signal": s})

		// Close the gRPC server
		server.GracefulStop()
		log.Println("Caught signal: ", s.String())
		// Exit the application with a 0 (success) status code.
		os.Exit(0)
	}()

	// Start the server 🥳
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err)
	}

	log.Println("Server gracefully stopped")
	return
}

func openDB(cfg model.Config) (*sql.DB, error) {
	log.Println("Connecting to database... 🚀")
	defer log.Println("Connected to database! 🥳")
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.DB.MaxOpenConn)
	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.DB.MaxIdleConn)
	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	// Set the maximum idle timeout.
	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func accessibleRoles() map[string][]string {
	//const authServicePath = "/auth.AuthService/"
	return map[string][]string{
		//authServicePath + "Register": {"admin"},
	}
}
