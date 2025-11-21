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
)

type AuthRepo interface {
	CreateUser(c context.Context, data models.Authentication) (res int64, err error)
	CreateDriver(c context.Context, data models.Authentication) (res int64, err error)
	LoginUser(c context.Context, data dto.UserLoginReq) (res models.Authentication, err error)
}

type AuthRepoImpl struct {
	db *gorm.DB
}

func (a *AuthRepoImpl) LoginUser(c context.Context, data dto.UserLoginReq) (res models.Authentication, err error) {
	if err := a.db.WithContext(c).First(&res, "email = ?", data.Email).Error; err != nil {
		return res, helper.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password)); err != nil {
		return res, helper.ErrPasswordIncorrect
	}

	return res, nil
}

func (a *AuthRepoImpl) CreateUser(c context.Context, data models.Authentication) (res int64, err error) {
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
			return 0, helper.ErrDuplicateEntry
		}

		return 0, helper.ErrDatabase
	}

	tx.Commit()

	return data.ID, nil
}

func (a *AuthRepoImpl) CreateDriver(c context.Context, data models.Authentication) (res int64, err error) {
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
			return 0, helper.ErrDuplicateEntry
		}

		return 0, helper.ErrDatabase
	}

	tx.Commit()

	return data.ID, nil
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &AuthRepoImpl{
		db: db,
	}
}
