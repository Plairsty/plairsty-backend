package main

import (
	"awesomeProject/internal/application"
	"awesomeProject/internal/interfaces/interceptors"
	sys "awesomeProject/internal/proto/health"
	studentPb "awesomeProject/internal/proto/student"
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
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := model.NewApplication(cfg, logger)

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
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection successfully opened")
	return db, nil
}
