package gRPC

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/pb"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
)

type GRPC struct {
	pb.UnimplementedDashboardServiceServer
	repo repository.DashboardRepo
}

func NewGRPC(repo repository.DashboardRepo) *GRPC {
	return &GRPC{
		repo: repo,
	}
}

func (a *GRPC) CreateOwner(ctx context.Context, data *pb.CreateOwnerReq) (*pb.CreateOwnerReq, error) {
	saveDir := "./uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	fullPath := data.Id + "_" + timestamp + data.Filename
	filePath := filepath.Join(saveDir, fullPath)
	if err := os.WriteFile(filePath, data.ProfilePicture, 0644); err != nil {
		return nil, err
	}

	res, err := a.repo.CreateBusinessOwner(ctx, models.OwnerDetails{
		ID:             data.Id,
		Name:           data.Name,
		Email:          data.Email,
		PhoneNumber:    data.PhoneNumber,
		NIK:            data.Nik,
		ProfilePicture: filePath,
		Verified:       false,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateOwnerReq{
		Id:          res.ID,
		Name:        res.Name,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Nik:         res.NIK,
	}, nil
}

func (a *GRPC) CreateGov(ctx context.Context, data *pb.CreateGovReq) (*pb.CreateGovReq, error) {
	img, format, err := image.Decode(bytes.NewReader(data.ProfilePicture))
	if err != nil {
		return nil, err
	}

	saveDir := "./uploads"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	fullPath := data.Id + "_" + timestamp
	filePath := filepath.Join(saveDir, fullPath)

	outFile, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer outFile.Close()

	switch format {
	case "jpeg":
		err = jpeg.Encode(outFile, img, nil)
	case "png":
		err = png.Encode(outFile, img)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err := os.WriteFile(filePath, data.ProfilePicture, 0644); err != nil {
		return nil, err
	}

	res, err := a.repo.CreateGovernment(ctx, models.GovDetails{
		ID:             data.Id,
		Name:           data.Name,
		Email:          data.Email,
		NIP:            data.Nip,
		ProfilePicture: filePath,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateGovReq{
		Id:    res.ID,
		Name:  res.Name,
		Email: res.Email,
		Nip:   res.NIP,
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
