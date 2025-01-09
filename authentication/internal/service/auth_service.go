package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type AuthService interface {
	CreateUserService(c context.Context, data dto.UserRegistrationsReq, role string) (res string, err *helper.ErrorStruct)
	CreateDriverService(c context.Context, data dto.DriverRegistrationsReq, role string) (res string, err *helper.ErrorStruct)
	CreateOwnerService(c context.Context, data dto.OwnerRegistrationsReq, role string) (res string, err *helper.ErrorStruct)
	LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
	SendResetPasswordService(c context.Context, email dto.ForgotPasswordReq) (res string, err *helper.ErrorStruct)
	ResetPassword(c context.Context, data dto.ResetPasswordReq, code string) (res string, err *helper.ErrorStruct)
	ChangePasswordService(c context.Context, id string, data dto.ChangePasswordReq) (res string, err *helper.ErrorStruct)
}

type AuthServiceImpl struct {
	AuthRepo repository.AuthRepo
}

func (a *AuthServiceImpl) CreateOwnerService(c context.Context, data dto.OwnerRegistrationsReq, role string) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	pw, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return "", err
	}

	resRepo, errRepo := a.AuthRepo.CreateUser(c, models.User{
		ID:       uuid.New().String(),
		Email:    data.Email,
		Password: string(pw),
		Role:     role,
	})

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrDuplicateEntry):
			code = fiber.StatusConflict
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) CreateDriverService(c context.Context, data dto.DriverRegistrationsReq, role string) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	pw, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return "", err
	}

	resRepo, errRepo := a.AuthRepo.CreateUser(c, models.User{
		ID:       uuid.New().String(),
		Email:    data.Email,
		Password: string(pw),
		Role:     role,
	})

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrDuplicateEntry):
			code = fiber.StatusConflict
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) ChangePasswordService(c context.Context, id string, data dto.ChangePasswordReq) (res string, err *helper.ErrorStruct) {
	hashedNewPassword, errHashed := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)

	if errHashed != nil {
		return res, &helper.ErrorStruct{
			Err:  errHashed,
			Code: fiber.StatusInternalServerError,
		}
	}

	resRepo, errRepo := a.AuthRepo.ChangePassword(c, data.OldPassword, string(hashedNewPassword), id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = fiber.StatusNotFound
		case errors.Is(errRepo, helper.ErrPasswordIncorrect):
			code = fiber.StatusUnauthorized
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) CreateUserService(c context.Context, data dto.UserRegistrationsReq, role string) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	pw, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return "", err
	}

	resRepo, errRepo := a.AuthRepo.CreateUser(c, models.User{
		ID:       uuid.New().String(),
		Email:    data.Email,
		Password: string(pw),
		Role:     role,
	})

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrDuplicateEntry):
			code = fiber.StatusConflict
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  helper.ErrBadRequest,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.AuthRepo.LoginUser(c, data)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = fiber.StatusNotFound
		case errors.Is(errRepo, helper.ErrPasswordIncorrect):
			code = fiber.StatusUnauthorized
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return dto.UserRegistrationsResp{
		ID:    resRepo.ID,
		Email: resRepo.Email,
		Role:  resRepo.Role,
	}, nil
}

func (a *AuthServiceImpl) SendResetPasswordService(c context.Context, email dto.ForgotPasswordReq) (res string, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(email); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	code := uuid.NewString()

	resRepo, errRepo := a.AuthRepo.SendResetPassword(c, email.Email, code)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = fiber.StatusNotFound
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
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
		"tiuq dxsj ubgf ztxf",
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	return "Link reset password telah dikirim ke email anda!", nil
}

func (a *AuthServiceImpl) ResetPassword(c context.Context, data dto.ResetPasswordReq, code string) (res string, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return "", &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	resRepo, errRepo := a.AuthRepo.ResetPassword(c, string(password), code)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = fiber.StatusNotFound
		default:
			code = fiber.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return resRepo, nil
}

func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return &AuthServiceImpl{
		AuthRepo: authRepo,
	}
}
