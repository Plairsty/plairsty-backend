package main

import (
	"flag"
	"log"
	"net"
	"os"
)

const version = "1.0.0"

type config struct {
	port string
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":4000", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	_ = &application{
		config: cfg,
		logger: logger,
	}
	lis, err := net.Listen("tcp", cfg.port)
	if err != nil {
		logger.Fatal(err)
	}
	log.Println("Starting server on port", cfg.port)
	defer func(conn net.Listener) {
		log.Println("Closing server on port", cfg.port)
		err := conn.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(lis)
}
