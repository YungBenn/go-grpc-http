package utils

import (
	"go-grpc-http/internal/pb"
	"go-grpc-http/internal/postgresql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCar(car postgresql.Car) *pb.Car {
	return &pb.Car{
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
}
