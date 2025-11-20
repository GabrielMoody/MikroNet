package controller

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/GabrielMoody/mikronet-driver-service/internal/dto"
	"github.com/GabrielMoody/mikronet-driver-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-driver-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

type DriverController interface {
	GetDriver(c *fiber.Ctx) error
	EditDriver(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	GetTripHistories(c *fiber.Ctx) error
	GetImage(c *fiber.Ctx) error
	GetAllDriverLastSeen(c *fiber.Ctx) error
	SetDriverLastSeen(c *fiber.Ctx) error
	GetQrisData(c *fiber.Ctx) error
}

type DriverControllerImpl struct {
	service service.DriverService
}

func readImage(image *multipart.FileHeader) ([]byte, error) {
	f, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileData, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func (a *DriverControllerImpl) GetQrisData(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetQrisData(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	code, errQr := qrcode.Encode(res, qrcode.Medium, 256)

	if errQr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": errQr,
		})
	}

	c.Response().Header.Set("Content-Type", "image/png")

	return c.Status(fiber.StatusOK).Send(code)
}

func (a *DriverControllerImpl) GetAllDriverLastSeen(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.service.GetAllLastSeen(ctx)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) SetDriverLastSeen(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	_, err := a.service.SetLastSeen(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
	})
}

func (a *DriverControllerImpl) GetImage(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.service.GetImage(ctx, id)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	img, errI := os.ReadFile(res)

	if errI != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": errI,
		})
	}

	ext := http.DetectContentType(img)

	c.Response().Header.Set("Content-Type", ext)

	return c.Status(fiber.StatusOK).Send(img)
}

func (a *DriverControllerImpl) GetDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetDriverDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"id":             res.ID,
			"name":           res.Name,
			"email":          res.Email,
			"license_number": res.LicenseNumber,
			"SIM":            res.SIM,
			"image":          os.Getenv("BASE_URL") + "/api/driver/images/" + res.ID,
		},
	})
}

func (a *DriverControllerImpl) EditDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	var data dto.EditDriverReq
	var fileDataPP []byte

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	image, err := c.FormFile("profile_picture")

	if err != nil || err == http.ErrMissingFile {
		log.Println("No Image uploaded!")
	} else {
		fileDataPP, err = readImage(image)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"errors": err.Error(),
			})
		}
	}

	_, errService := a.service.EditDriverDetails(ctx, payload["id"].(string), data, fileDataPP)

	if errService != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Data berhasil diperbarui!",
	})
}

func (a *DriverControllerImpl) GetTripHistories(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetTripHistories(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) GetStatus(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetStatus(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"status": res,
		},
	})
}

func (a *DriverControllerImpl) SetStatus(c *fiber.Ctx) error {
	var data dto.StatusReq
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	res, err := a.service.SetStatus(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"status": res,
		},
	})
}

func NewDriverController(service service.DriverService) DriverController {
	return &DriverControllerImpl{
		service: service,
	}
}
