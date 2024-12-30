package repository

import (
	"context"
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	CreateUser(c context.Context, data models.User) (res string, err error)
	LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error)
	SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error)
	ResetPassword(c context.Context, password string, code string) (res string, err error)
	ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error)
}

type AuthRepoImpl struct {
	db *gorm.DB
}

func (a *AuthRepoImpl) ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "id = ?", id).Error; err != nil {
		return "Email not found!", gorm.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return "Incorrect password!", err
	}

	if err := a.db.WithContext(c).Model(&user).Update("password", newPassword).Error; err != nil {
		return "Error while Updating data", err
	}

	return "Success updating the password", nil
}

func (a *AuthRepoImpl) LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "email = ?", data.Email).Error; err != nil {
		return res, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return res, err
	}

	return user, nil
}

func (a *AuthRepoImpl) CreateUser(c context.Context, data models.User) (res string, err error) {
	result := a.db.WithContext(c).Create(&data)

	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (a *AuthRepoImpl) SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "email = ?", email).Error; err != nil {
		return data, gorm.ErrRecordNotFound
	}

	rp := models.ResetPassword{
		UserID: user.ID,
		Code:   code,
	}

	if err := a.db.WithContext(c).Create(&rp).Error; err != nil {
		return data, err
	}

	return rp, nil
}

func (a *AuthRepoImpl) ResetPassword(c context.Context, password string, code string) (res string, err error) {
	var rp models.ResetPassword

	if err := a.db.WithContext(c).First(&rp, "code = ?", code).Error; err != nil {
		return "Link has expired or not valid", err
	}

	user := models.User{
		Password: password,
	}

	if err := a.db.WithContext(c).Where("id = ?", rp.UserID).Updates(&user).Error; err != nil {
		return "something happened", err
	}

	if err := a.db.WithContext(c).Delete(&rp).Error; err != nil {
		return "something happened", err
	}

	return "Success", nil
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}
