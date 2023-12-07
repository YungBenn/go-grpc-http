package main

import (
	"context"
	"fmt"
	"go-grpc-http/internal/http/grpc"
	"go-grpc-http/internal/postgresql"
	"go-grpc-http/pkg/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log := utils.LoggerInit()
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")	
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), 
		os.Getenv("DB_NAME"), 
		os.Getenv("DB_SSL_MODE"),
	)
	db, err := postgresql.Connect(context.Background(), log, dbURL)
	if err != nil {
		log.Panic(err)
	}

	serverUrl := fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	grpc.RunServer(log, db, serverUrl)
}
