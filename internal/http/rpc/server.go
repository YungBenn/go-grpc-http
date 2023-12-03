package rpc

import (
	"go-grpc-http/internal/pb"
	"go-grpc-http/internal/postgresql"
	"go-grpc-http/internal/service"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(log *logrus.Logger,db *pgxpool.Pool, address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	repo := postgresql.NewRepository(db)
	srv := service.NewCarServiceServer(log, repo)
	pb.RegisterCarServiceServer(server, srv)
	reflection.Register(server)

	err = server.Serve(listen)
	if err != nil {
		log.Errorf("cannot start gRPC server: %v", err)
	}
}