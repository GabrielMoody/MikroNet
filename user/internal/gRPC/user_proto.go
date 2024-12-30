package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/pb"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
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
	format := "02-01-2006"
	date, err := time.Parse(format, req.User.DateOfBirth)

	if err != nil {
		return res, err
	}

	data := model.UserDetails{
		ID:          req.User.Id,
		FirstName:   req.User.FirstName,
		LastName:    req.User.LastName,
		Email:       req.User.Email,
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

func (a *GRPC) GetUsers(ctx context.Context, _ *pb.Empty) (res *pb.Users, err error) {
	resRepo, err := a.repo.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	var users []*pb.User

	for _, v := range resRepo {
		users = append(users, &pb.User{
			Id:          v.ID,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			Age:         uint32(v.Age),
			Gender:      v.Gender,
		})
	}

	return &pb.Users{
		Users: users,
	}, nil
}

func (a *GRPC) GetUserDetails(ctx context.Context, req *pb.GetUserDetailsRequest) (res *pb.User, err error) {
	resRepo, err := a.repo.GetUserDetails(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
		Age:         uint32(resRepo.Age),
	}, nil
}
