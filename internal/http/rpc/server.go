package rpc

import (
	"context"
	"go-grpc-http/internal/pb"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(ctx context.Context, log *logrus.Logger, srv pb.CarServiceServer, port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterCarServiceServer(server, srv)

	err = server.Serve(listen)
	if err != nil {
		log.Errorf("cannot start gRPC server: %v", err)
	}
}