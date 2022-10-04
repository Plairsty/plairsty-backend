package main

import (
	"awesomeProject/internal/application"
	"awesomeProject/internal/interfaces/interceptors"
	sys "awesomeProject/internal/proto/health"
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
	
	// After all configuration is done, we can create a new application instance.
	app := model.NewApplication(cfg, logger, db)

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
		grpc.UnaryInterceptor(interceptors.Unary()),
		grpc.StreamInterceptor(interceptors.Stream()),
	)

	// Register your services here 🥵
	// If you found an error stating that certain method is not implemented,
	// Check if you have implemented the method in your service implementation
	// And then check if you have added unimplemented method in your service struct
	sys.RegisterHealthCheckServer(server, app)
	studentPb.RegisterStudentServiceServer(server, app)

	// register reflection service on gRPC server.
	reflection.Register(server)

	// Get service info
	for _, service := range server.GetServiceInfo() {
		go func(s grpc.ServiceInfo) {
			for _, method := range s.Methods {
				log.Println("Invoking Method:", method.Name)
			}
		}(service)
	}

	// Start the server 🥳
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err)
	}
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
