package repository

import (
	"context"
	"errors"
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	mysql2 "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ProfileRepo interface {
	GetUser(c context.Context, id string) (res models.User, err error)
	CreateUser(c context.Context, data models.User) (res string, err error)
	LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error)
	UpdateUser(c context.Context, id string, data models.User) (res string, err error)
	DeleteUser(c context.Context, id string) (res string, err error)
	ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error)
	ForgotPassword(c context.Context, id string) (res string, err error)
}

type ProfileRepoImpl struct {
	db *gorm.DB
}

func (a *ProfileRepoImpl) LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "email = ?", data.Email).Error; err != nil {
		return res, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return res, err
	}

	return user, nil
}

func (a *ProfileRepoImpl) GetUser(c context.Context, id string) (res models.User, err error) {
	if err := a.db.WithContext(c).First(&res).Where("id = ?", id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *ProfileRepoImpl) CreateUser(c context.Context, data models.User) (res string, err error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user := models.User{
		ID:           data.ID,
		NamaLengkap:  data.NamaLengkap,
		Email:        data.Email,
		NomorTelepon: data.NomorTelepon,
		Password:     string(pw),
	}

	result := a.db.WithContext(c).Create(&user)

	var mysqlErr *mysql2.MySQLError

	if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		return res, errors.New("email dan/atau nomor telepon telah terdaftar")
	}

	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
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
		return "Error while Updating data", gorm.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(user.Password)); err != nil {
		return "Incorrect password", err
	}

	if err := a.db.WithContext(c).Model(&models.User{}).Update("password", newPassword).Error; err != nil {
		return "Error while Updating data", err
	}

	return res, nil
}

func (a *ProfileRepoImpl) ForgotPassword(c context.Context, id string) (res string, err error) {
	//TODO implement me
	panic("implement me")
}

func NewProfileRepo(db *gorm.DB) ProfileRepo {
	return &ProfileRepoImpl{
		db: db,
	}
}
