package service

import (
	"context"
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"os"
	"time"
)

type ProfileService interface {
	CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct)
	CreateDriverService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct)
	LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
	SendResetPasswordService(c context.Context, email dto.ForgotPasswordReq) (res string, err *helper.ErrorStruct)
	ResetPassword(c context.Context, data dto.ResetPasswordReq, code string) (res string, err *helper.ErrorStruct)
}

type ProfileServiceImpl struct {
	ProfileRepo repository.ProfileRepo
}

func (a *ProfileServiceImpl) CreateDriverService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct) {
	//TODO implement me
	panic("implement me")
}

func (a *ProfileServiceImpl) CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	format := "01-02-2006"
	date, _ := time.Parse(format, data.DateOfBirth)

	pw, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return "", err
	}

	resRepo, errRepo := a.ProfileRepo.CreateUser(c, models.User{
		ID:          uuid.New().String(),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Password:    string(pw),
		Age:         data.Age,
		DateOfBirth: &date,
		Role:        "user",
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *ProfileServiceImpl) LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProfileRepo.LoginUser(c, data)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return dto.UserRegistrationsResp{
		ID:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
		Role:        resRepo.Role,
	}, nil
}

func (a *ProfileServiceImpl) SendResetPasswordService(c context.Context, email dto.ForgotPasswordReq) (res string, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(email); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	code := uuid.NewString()

	resRepo, errRepo := a.ProfileRepo.SendResetPassword(c, email.Email, code)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "Mikronet <test.mikronet@gmail.com>"
	const CONFIG_AUTH_EMAIL = "test.mikronet@gmail.com"

	html := fmt.Sprintf(`
		<a href="http://localhost:8000/auth/api/auth/reset-password/%s"
        style="
		color: #fff;
		background-color: #0069d9;
		display: inline-block;
        font-weight: 400;
        text-align: center;
        white-space: nowrap;
        vertical-align: middle;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        border: 1px solid transparent;
        padding: .375rem .75rem;
        font-size: 1rem;
        line-height: 1.5;
        border-radius: .25rem;
		text-decoration: none;">Reset Password</a>
	`, resRepo.Code)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email.Email)
	mailer.SetHeader("Subject", "Reset Password")
	mailer.SetAddressHeader("Cc", CONFIG_AUTH_EMAIL, "Mikronet <test.mikronet@gmail.com>")
	mailer.SetBody("text/html", html)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		os.Getenv("EMAIL_PASSWORD"),
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	return "Link reset password telah dikirim ke email anda!", nil
}

func (a *ProfileServiceImpl) ResetPassword(c context.Context, data dto.ResetPasswordReq, code string) (res string, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return "", &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	resRepo, errRepo := a.ProfileRepo.ResetPassword(c, string(password), code)

	if errRepo != nil {
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewProfileService(ProfileRepo repository.ProfileRepo) ProfileService {
	return &ProfileServiceImpl{
		ProfileRepo: ProfileRepo,
	}
}
