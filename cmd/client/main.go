package main

import (
	"context"
	"fmt"
	"go-grpc-http/internal/pb"
	"go-grpc-http/pkg/utils"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log := utils.LoggerInit()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")	
	}

	serverUrl := fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	conn, err := grpc.Dial(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()

	err = pb.RegisterCarServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	clientUrl := fmt.Sprintf("%v:%v", os.Getenv("CLIENT_HOST"), os.Getenv("CLIENT_PORT"))
	gwServer := &http.Server{
		Addr: clientUrl,
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
}