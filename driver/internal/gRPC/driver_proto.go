package gRPC

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/GabrielMoody/mikronet-driver-service/internal/model"
	"github.com/GabrielMoody/mikronet-driver-service/internal/pb"
	"github.com/GabrielMoody/mikronet-driver-service/internal/repository"
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

func (a *GRPC) GetDrivers(ctx context.Context, req *pb.ReqDrivers) (res *pb.Drivers, err error) {
	resRepo, err := a.repo.GetAllDrivers(ctx, &req.Verified)

	if err != nil {
		return nil, err
	}

	var drivers []*pb.Driver

	for _, v := range resRepo {
		drivers = append(drivers, &pb.Driver{
			Id:            v.ID,
			Name:          v.Name,
			Email:         v.Email,
			PhoneNumber:   v.PhoneNumber,
			LicenseNumber: v.LicenseNumber,
			Sim:           v.SIM,
			ImageUrl:      os.Getenv("BASE_URL") + "/api/driver/images/" + v.ID,
		})
	}

	return &pb.Drivers{
		Drivers: drivers,
	}, nil
}

func (a *GRPC) GetDriverDetails(ctx context.Context, data *pb.ReqByID) (res *pb.Driver, err error) {
	resRepo, err := a.repo.GetDriverDetails(ctx, data.Id)

	if err != nil {
		return nil, err
	}

	var imageurl string

	if resRepo.ProfilePicture != "" {
		imageurl = os.Getenv("BASE_URL") + "/api/driver/images/" + resRepo.ID
	}

	return &pb.Driver{
		Id:            resRepo.ID,
		Name:          resRepo.Name,
		Email:         resRepo.Email,
		PhoneNumber:   resRepo.PhoneNumber,
		LicenseNumber: resRepo.LicenseNumber,
		Sim:           resRepo.SIM,
		ImageUrl:      imageurl,
	}, nil
}

func (a *GRPC) CreateDriver(ctx context.Context, data *pb.CreateDriverRequest) (res *pb.Driver, err error) {
	saveDir := "./uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	fullPath := data.Id + "_" + timestamp
	filePath := filepath.Join(saveDir, fullPath)
	if err := os.WriteFile(filePath, data.ProfilePicture, 0644); err != nil {
		return nil, err
	}

	resRepo, err := a.repo.CreateDriver(ctx, model.DriverDetails{
		ID:             data.Id,
		Name:           data.Name,
		Email:          data.Email,
		PhoneNumber:    data.PhoneNumber,
		LicenseNumber:  data.LicenseNumber,
		SIM:            data.Sim,
		ProfilePicture: filePath,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Driver{
		Id:            resRepo.ID,
		Name:          resRepo.Name,
		Email:         resRepo.Email,
		PhoneNumber:   resRepo.PhoneNumber,
		LicenseNumber: resRepo.LicenseNumber,
		Sim:           resRepo.SIM,
	}, nil
}

func (a *GRPC) SetStatusVerified(ctx context.Context, data *pb.ReqByID) (res *pb.Driver, err error) {
	resRepo, err := a.repo.SetVerified(ctx, model.DriverDetails{
		ID:       data.Id,
		Verified: true,
	})

	if err != nil {
		return nil, err
	}

	return &pb.Driver{
		Id:       resRepo.ID,
		Verified: resRepo.Verified,
	}, nil
}
