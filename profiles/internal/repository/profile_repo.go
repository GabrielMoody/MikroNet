package repository

import (
	"context"
	"errors"
	"github.com/GabrielMoody/mikroNet/profiles/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ProfileRepo interface {
	GetUser(c context.Context, id string) (res models.User, err error)
	UpdateUser(c context.Context, id string, data models.User) (res string, err error)
	DeleteUser(c context.Context, id string) (res string, err error)
	ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error)
	ForgotPassword(c context.Context, id string, newPassword string) (res string, err error)
}

type ProfileRepoImpl struct {
	db *gorm.DB
}

func (a *ProfileRepoImpl) GetUser(c context.Context, id string) (res models.User, err error) {
	if err := a.db.WithContext(c).First(&res).Where("id = ?", id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *ProfileRepoImpl) UpdateUser(c context.Context, id string, data models.User) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, id).Error; err != nil {
		return "Error while Updating data", gorm.ErrRecordNotFound
	}

	if err := a.db.WithContext(c).Model(&user).Updates(data).Error; err != nil {
		return "Error while Updating data", err
	}

	return res, nil
}

func (a *ProfileRepoImpl) DeleteUser(c context.Context, id string) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user).Where("id = ?", id).Error; err != nil {
		return res, errors.New("data pengguna tidak ditemukan")
	}

	if err := a.db.WithContext(c).Delete(&user).Error; err != nil {
		return res, err
	}

	return "berhasil menghapus profile", nil
}

func (a *ProfileRepoImpl) ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).Model(&user).Where("id = ?", id).Error; err != nil {
		return "Email not found!", gorm.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(user.Password)); err != nil {
		return "Incorrect password!", err
	}

	if err := a.db.WithContext(c).Model(&models.User{}).Update("password", newPassword).Error; err != nil {
		return "Error while Updating data", err
	}

	return res, nil
}

func (a *ProfileRepoImpl) ForgotPassword(c context.Context, id string, newPassword string) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).Model(&user).Where("id = ?", id).Error; err != nil {
		return "Email not found!", gorm.ErrRecordNotFound
	}

	if err := a.db.WithContext(c).Model(&models.User{}).Update("password", newPassword).Error; err != nil {
		return "Error while Updating data", err
	}

	return res, nil
}

func NewProfileRepo(db *gorm.DB) ProfileRepo {
	return &ProfileRepoImpl{
		db: db,
	}
}
