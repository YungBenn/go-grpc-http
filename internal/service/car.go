package service

import (
	"context"
	"go-grpc-http/internal/pb"
	"go-grpc-http/internal/postgresql"
	"go-grpc-http/pkg/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CarServiceServer struct {
	pb.UnimplementedCarServiceServer
	log  *logrus.Logger
	repo postgresql.Repository
}

// CreateCar implements pb.CarServiceServer.
func (s *CarServiceServer) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarResponse, error) {
	arg := postgresql.CreateCarParams{
		Name:        req.Name,
		Model:       req.Model,
		Color:       req.Color,
		Year:        req.Year,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	car, err := s.repo.CreateCar(ctx, arg)
	if err != nil {
		s.log.Error("Error creating car: ", err)
		return nil, err
	}

	s.log.Info("Car created: ", car.ID)
	return &pb.CreateCarResponse{
		Car: utils.ConvertCar(car),
	}, nil
}

// DeleteCar implements pb.CarServiceServer.
func (s *CarServiceServer) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteCarResponse, error) {
	carID := uuid.MustParse(req.Id)
	err := s.repo.DeleteCar(ctx, carID)
	if err != nil {
		s.log.Error("Error deleting car: ", err)
		return nil, err
	}

	s.log.Info("Car deleted: ", req.Id)
	return &pb.DeleteCarResponse{
		Id: req.Id,
	}, nil
}

// ListCar implements pb.CarServiceServer.
func (s *CarServiceServer) ListCar(ctx context.Context, req *pb.ListCarRequest) (*pb.ListCarResponse, error) {
	limit, err := strconv.ParseInt(req.GetLimit(), 10, 32)
	if err != nil {
		s.log.Error("Error parsing limit: ", err)
		return nil, err
	}

	cars, err := s.repo.ListCars(ctx, int32(limit))
	if err != nil {
		s.log.Error("Error listing cars: ", err)
		return nil, err
	}

	carList := []*pb.Car{}
	for _, car := range cars {
		carPB := &pb.Car{
			Id:          car.ID.String(),
			Name:        car.Name,
			Model:       car.Model,
			Color:       car.Color,
			Year:        car.Year,
			Price:       car.Price,
			Image:       car.Image,
			Description: car.Description,
			CreatedAt:   timestamppb.New(car.CreatedAt),
			UpdatedAt:   timestamppb.New(car.UpdatedAt),
		}
		carList = append(carList, carPB)
	}

	s.log.Info("Listing cars susccessful")
	return &pb.ListCarResponse{
		Cars: carList,
	}, nil
}

// ReadCar implements pb.CarServiceServer.
func (s *CarServiceServer) ReadCar(ctx context.Context, req *pb.ReadCarRequest) (*pb.ReadCarResponse, error) {
	carID := uuid.MustParse(req.Id)
	car, err := s.repo.GetCar(ctx, carID)
	if err != nil {
		s.log.Error("Error reading car: ", err)
		return nil, err
	}

	return &pb.ReadCarResponse{
		Car: utils.ConvertCar(car),
	}, nil
}

// UpdateCar implements pb.CarServiceServer.
func (s *CarServiceServer) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.UpdateCarResponse, error) {
	arg := postgresql.UpdateCarParams{
		ID:          uuid.MustParse(req.Id),
		Name:        req.Name,
		Model:       req.Model,
		Color:       req.Color,
		Year:        req.Year,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
		UpdatedAt:   time.Now(),
	}

	car, err := s.repo.UpdateCar(ctx, arg)
	if err != nil {
		s.log.Error("Error updating car: ", err)
		return nil, err
	}

	s.log.Info("Car updated: ", req.Id)
	return &pb.UpdateCarResponse{
		Car: utils.ConvertCar(car),
	}, nil
}

func NewCarServiceServer(log *logrus.Logger, repo postgresql.Repository) pb.CarServiceServer {
	return &CarServiceServer{
		UnimplementedCarServiceServer: pb.UnimplementedCarServiceServer{},
		log:                           log,
		repo:                          repo,
	}
}
