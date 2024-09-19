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
	CreateUser(c context.Context, data models.User) (res string, err error)
	CreateDriver(c context.Context, user models.User, driver models.Driver) (res string, err error)
	LoginUser(c context.Context, data dto.UserLoginReq) (res models.User, err error)
	SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error)
	ResetPassword(c context.Context, password string, code string) (res string, err error)
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

func (a *ProfileRepoImpl) CreateUser(c context.Context, data models.User) (res string, err error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	//user := models.User{
	//	ID:          data.ID,
	//	FirstName:   data.FirstName,
	//	LastName:    data.LastName,
	//	Email:       data.Email,
	//	PhoneNumber: data.PhoneNumber,
	//	Password:    string(pw),
	//	Age:         data.Age,
	//	DateOfBirth: data.DateOfBirth,
	//	Gender:      data.Gender,
	//	Role:        "user",
	//}

	data.Password = string(pw)
	data.Role = "user"

	result := a.db.WithContext(c).Create(&data)

	var mysqlErr *mysql2.MySQLError

	if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		return res, errors.New("email or phone number already exist")
	}

	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (a *ProfileRepoImpl) CreateDriver(c context.Context, user models.User, driver models.Driver) (res string, err error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	//data := models.User{
	//	ID:          user.ID,
	//	FirstName:   user.FirstName,
	//	LastName:    user.LastName,
	//	Email:       user.Email,
	//	PhoneNumber: user.PhoneNumber,
	//	Password:    string(pw),
	//	Age:         user.Age,
	//	DateOfBirth: user.DateOfBirth,
	//	Gender:      user.Gender,
	//	Role:        "driver",
	//}

	user.Password = string(pw)
	user.Role = "driver"

	result := a.db.WithContext(c).Create(&user)

	var mysqlErr *mysql2.MySQLError

	if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
		return res, errors.New("email or phone number already exist")
	}

	if result.Error != nil {
		return res, result.Error
	}

	return user.ID, nil
}

func (a *ProfileRepoImpl) SendResetPassword(c context.Context, email string, code string) (data models.ResetPassword, err error) {
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

func (a *ProfileRepoImpl) ResetPassword(c context.Context, password string, code string) (res string, err error) {
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

func NewProfileRepo(db *gorm.DB) ProfileRepo {
	return &ProfileRepoImpl{
		db: db,
	}
}
