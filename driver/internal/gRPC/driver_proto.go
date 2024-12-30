package gRPC

import (
	"context"
	"github.com/GabrielMoody/mikroNet/driver/internal/model"
	"github.com/GabrielMoody/mikroNet/driver/internal/pb"
	"github.com/GabrielMoody/mikroNet/driver/internal/repository"
)

type GRPC struct {
	pb.UnimplementedDriverServiceServer
	repo repository.DriverRepo
}

func NewgRPC(repo repository.DriverRepo) *GRPC {
	return &GRPC{
		repo: repo,
	}
}

func (a *GRPC) GetDrivers(ctx context.Context, _ *pb.Empty) (res *pb.Drivers, err error) {
	resRepo, err := a.repo.GetAllDrivers(ctx)

	if err != nil {
		return nil, err
	}

	var drivers []*pb.Driver

	for _, v := range resRepo {
		drivers = append(drivers, &pb.Driver{
			Id:          v.ID,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			Age:         uint32(v.Age),
		})
	}

	return &pb.Drivers{
		Drivers: drivers,
	}, nil
}

func (a *GRPC) GetDriverDetails(ctx context.Context, data *pb.ReqDriverDetails) (res *pb.Driver, err error) {
	resRepo, err := a.repo.GetDriverDetails(ctx, data.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Driver{
		Id:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
		Age:         uint32(resRepo.Age),
	}, nil
}

func (a *GRPC) CreateDriver(ctx context.Context, data *pb.CreateDriverRequest) (res *pb.Driver, err error) {
	resRepo, err := a.repo.CreateDriver(ctx, model.DriverDetails{
		ID:            data.Id,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		Email:         data.Email,
		PhoneNumber:   data.PhoneNumber,
		Age:           int32(data.Age),
		LicenseNumber: data.LicenseNumber,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Driver{
		Id:                 resRepo.ID,
		FirstName:          resRepo.FirstName,
		LastName:           resRepo.LastName,
		Email:              resRepo.Email,
		PhoneNumber:        resRepo.PhoneNumber,
		Age:                uint32(resRepo.Age),
		RegistrationNumber: resRepo.LicenseNumber,
	}, nil
}
