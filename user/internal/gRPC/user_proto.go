package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/pb"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type GRPC struct {
	pb.UnimplementedUserServiceServer
	repo repository.UserRepo
}

func NewgRPC(repo repository.UserRepo) *GRPC {
	return &GRPC{
		repo: repo,
	}
}

func (a *GRPC) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (res *pb.CreateUserResponse, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)

	if err != nil {
		return res, err
	}

	format := "02-01-2006"
	date, err := time.Parse(format, req.User.DateOfBirth)

	if err != nil {
		return res, err
	}

	data := model.User{
		ID:          req.User.Id,
		FirstName:   req.User.FirstName,
		LastName:    req.User.LastName,
		Email:       req.User.Email,
		Password:    string(hashed),
		PhoneNumber: req.User.PhoneNumber,
		Age:         int32(req.User.Age),
		Gender:      req.User.Gender,
		DateOfBirth: date,
	}

	resRepo, err := a.repo.CreateUser(ctx, data)

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: resRepo.ID,
	}, nil
}
