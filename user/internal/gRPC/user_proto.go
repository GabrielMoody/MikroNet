package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/pb"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"os"
	"path/filepath"
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
	saveDir := "./uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	fullPath := req.User.Id + "_" + timestamp + req.User.Filename
	filePath := filepath.Join(saveDir, fullPath)

	if err := os.WriteFile(filePath, req.User.ProfilePicture, 0644); err != nil {
		return nil, err
	}

	data := model.UserDetails{
		ID:             req.User.Id,
		FirstName:      req.User.FirstName,
		LastName:       req.User.LastName,
		Email:          req.User.Email,
		PhoneNumber:    req.User.PhoneNumber,
		ProfilePicture: filePath,
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

	return &pb.User{
		Id:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
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
