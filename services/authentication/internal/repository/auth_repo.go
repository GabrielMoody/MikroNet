package repository

import (
	"context"
	"errors"


	"github.com/GabrielMoody/MikroNet/services/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mattn/go-sqlite3"

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
	if err := a.db.WithContext(c).First(&res, "username = ?", data.Email).Error; err != nil {
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

		// Check for PostgreSQL unique constraint error
		var psqlErr *pgconn.PgError
		if errors.As(err, &psqlErr) && psqlErr.Code == "23505" {
			return 0, helper.ErrDuplicateEntry
		}

		// Check for GORM's duplicated key error (should be database agnostic, but might not catch all cases)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, helper.ErrDuplicateEntry
		}

		// Check for SQLite specific unique constraint error
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
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

		// Check for PostgreSQL unique constraint error
		var psqlErr *pgconn.PgError
		if errors.As(err, &psqlErr) && psqlErr.Code == "23505" {
			return 0, helper.ErrDuplicateEntry
		}

		// Check for GORM's duplicated key error (should be database agnostic, but might not catch all cases)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, helper.ErrDuplicateEntry
		}

		// Check for SQLite specific unique constraint error
		var sqliteErr sqlite3.Error
					if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {			return 0, helper.ErrDuplicateEntry
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