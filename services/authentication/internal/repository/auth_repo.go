package repository

import (
	"context"
	"errors"

	"github.com/GabrielMoody/mikronet-auth-service/internal/dto"
	"github.com/GabrielMoody/mikronet-auth-service/internal/helper"
	"github.com/GabrielMoody/mikronet-auth-service/internal/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthRepo interface {
	CreateUser(c context.Context, data models.User) (res string, err error)
	CreateDriver(c context.Context, data models.User) (res string, err error)
	LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error)
	SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error)
	ResetPassword(c context.Context, password string, code string) (res string, err error)
	ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error)
	DeleteUser(c context.Context, id string) (res models.User, err error)
	isBlocked(c context.Context, id string) (bool, error)
	isVerified(c context.Context, id string) (bool, error)
}

type AuthRepoImpl struct {
	db *gorm.DB
}

func (a *AuthRepoImpl) isBlocked(c context.Context, id string) (bool, error) {
	var res models.BlockedAccount
	if err := a.db.WithContext(c).First(&res, "account_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, helper.ErrDatabase
	}

	return true, nil
}

func (a *AuthRepoImpl) isVerified(c context.Context, id string) (bool, error) {
	var res models.DriverDetails
	if err := a.db.WithContext(c).First(&res, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, helper.ErrDatabase
	}

	return res.Verified, nil
}

func (a *AuthRepoImpl) DeleteUser(c context.Context, id string) (res models.User, err error) {
	if err := a.db.WithContext(c).Delete(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *AuthRepoImpl) ChangePassword(c context.Context, oldPassword, newPassword, id string) (res string, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "id = ?", id).Error; err != nil {
		return "Email not found!", helper.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return "Incorrect password!", helper.ErrPasswordIncorrect
	}

	if err := a.db.WithContext(c).Model(&user).Update("password", newPassword).Error; err != nil {
		return "Error while Updating data", helper.ErrDatabase
	}

	return "Success updating the password", nil
}

func (a *AuthRepoImpl) LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error) {
	if err := a.db.WithContext(c).First(&res, "email = ?", data.Email).Error; err != nil {
		return res, helper.ErrNotFound
	}

	b, _ := a.isBlocked(c, res.ID)
	if b {
		return res, helper.ErrBlockedAccount
	}

	if res.Role == "driver" {
		v, _ := a.isVerified(c, res.ID)
		if !v {
			return res, helper.ErrNotVerified
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password)); err != nil {
		return res, helper.ErrPasswordIncorrect
	}

	return res, nil
}

func (a *AuthRepoImpl) CreateUser(c context.Context, data models.User) (res string, err error) {
	tx := a.db.WithContext(c).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()

		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return "", helper.ErrDuplicateEntry
		}

		return "", helper.ErrDatabase
	}

	tx.Commit()

	return data.ID, nil
}

func (a *AuthRepoImpl) CreateDriver(c context.Context, data models.User) (res string, err error) {
	tx := a.db.WithContext(c).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()

		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return "", helper.ErrDuplicateEntry
		}

		return "", helper.ErrDatabase
	}

	tx.Commit()

	return data.ID, nil
}

func (a *AuthRepoImpl) SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error) {
	var user models.User

	if err := a.db.WithContext(c).First(&user, "email = ?", email).Error; err != nil {
		return data, helper.ErrNotFound
	}

	rp := models.ResetPassword{
		UserID: user.ID,
		Code:   code,
	}

	if err := a.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"code"}),
	}).Create(&rp).Error; err != nil {
		return data, helper.ErrDatabase
	}

	return rp, nil
}

func (a *AuthRepoImpl) ResetPassword(c context.Context, password string, code string) (res string, err error) {
	var rp models.ResetPassword

	if err := a.db.WithContext(c).First(&rp, "code = ?", code).Error; err != nil {
		return "Link has expired or not valid", helper.ErrExpired
	}

	user := models.User{
		Password: password,
	}

	if err := a.db.WithContext(c).Where("id = ?", rp.UserID).Updates(&user).Error; err != nil {
		return "something happened", helper.ErrDatabase
	}

	if err := a.db.WithContext(c).Delete(&rp).Error; err != nil {
		return "something happened", helper.ErrDatabase
	}

	return "Success", nil
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}
