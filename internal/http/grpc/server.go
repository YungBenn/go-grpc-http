package grpc

import (
	"go-grpc-http/internal/pb"
	"go-grpc-http/internal/postgresql"
	"go-grpc-http/internal/service"
	"net"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var keep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, 
	PermitWithoutStream: true,            
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, 
	MaxConnectionAge:      30 * time.Second, 
	MaxConnectionAgeGrace: 5 * time.Second,  
	Time:                  5 * time.Second,  
	Timeout:               1 * time.Second,  
}

func RunServer(log *logrus.Logger, db *pgxpool.Pool, address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(keep), grpc.KeepaliveParams(kasp))
	repo := postgresql.NewRepository(db)
	srv := service.NewCarServiceServer(log, repo)
	pb.RegisterCarServiceServer(server, srv)
	reflection.Register(server)

	log.Info("starting gRPC server...")
	err = server.Serve(listen)
	if err != nil {
		log.Errorf("cannot start gRPC server: %v", err)
	}
}
