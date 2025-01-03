package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/pb"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"strconv"
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
		formattedDate := v.DateOfBirth.Format("02-01-2006")
		users = append(users, &pb.User{
			Id:          v.ID,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Email:       v.Email,
			DateOfBirth: formattedDate,
			PhoneNumber: v.PhoneNumber,
			Age:         uint32(v.Age),
			Gender:      v.Gender,
		})
	}

	return &pb.Users{
		Users: users,
	}, nil
}

func (a *GRPC) GetUserDetails(ctx context.Context, req *pb.GetByIDRequest) (res *pb.User, err error) {
	resRepo, err := a.repo.GetUserDetails(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	formattedDate := resRepo.DateOfBirth.Format("02-01-2006")

	return &pb.User{
		Id:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
		Age:         uint32(resRepo.Age),
		DateOfBirth: formattedDate,
	}, nil
}

func (a *GRPC) GetReviews(ctx context.Context, _ *pb.Empty) (res *pb.GetReviewsResponse, err error) {
	resRepo, err := a.repo.GetAllReviews(ctx)

	if err != nil {
		return nil, err
	}

	var reviews []*pb.Review

	for _, v := range resRepo {
		reviews = append(reviews, &pb.Review{
			Id:       strconv.Itoa(v.ID),
			UserId:   v.UserID,
			DriverId: v.DriverID,
			Comment:  v.Comment,
			Star:     uint32(v.Star),
		})
	}

	return &pb.GetReviewsResponse{
		Reviews: reviews,
	}, nil
}

func (a *GRPC) GetReviewsByID(ctx context.Context, req *pb.GetByIDRequest) (res *pb.Review, err error) {
	resRepo, err := a.repo.GetReviewsByID(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Review{
		Id:       strconv.Itoa(resRepo.ID),
		UserId:   resRepo.UserID,
		DriverId: resRepo.DriverID,
		Comment:  resRepo.Comment,
		Star:     uint32(resRepo.Star),
	}, nil
}
