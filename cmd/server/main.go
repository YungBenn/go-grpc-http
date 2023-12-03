package main

import (
	"context"
	"go-grpc-http/internal/http/rpc"
	"go-grpc-http/internal/postgresql"
	"go-grpc-http/internal/service"
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

	db, err := postgresql.Connect(context.Background(), log, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic(err)
	}

	repo := postgresql.NewRepository(db)
	srv := service.NewCarServiceServer(log, repo)

	rpc.RunServer(context.Background(), log, srv, os.Getenv("GRPC_PORT"))
}
