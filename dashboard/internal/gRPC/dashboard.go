package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/repository"
)

type GRPC struct {
	pb.UnimplementedOwnerServiceServer
	repo repository.DashboardRepo
}

func NewGRPC(repo repository.DashboardRepo) *GRPC {
	return &GRPC{
		repo: repo,
	}
}

func (a *GRPC) CreateOwner(ctx context.Context, data *pb.CreateOwnerReq) (*pb.CreateOwnerReq, error) {
	res, err := a.repo.CreateBusinessOwner(ctx, models.OwnerDetails{
		ID:          data.Id,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		NIK:         data.Nik,
		Verified:    false,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateOwnerReq{
		Id:          res.ID,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Nik:         res.NIK,
	}, nil
}

func (a *GRPC) IsBlocked(ctx context.Context, data *pb.IsBlockedReq) (*pb.IsBlockedRes, error) {
	res, err := a.repo.IsBlocked(ctx, data.Id)

	if err != nil {
		return nil, err
	}

	return &pb.IsBlockedRes{
		IsBlocked: res,
	}, nil
}
