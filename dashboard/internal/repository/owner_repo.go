package repository

//
//import (
//	"context"
//	"gorm.io/gorm"
//	"time"
//)
//
//type OwnerRepo interface {
//	RegisterBusinessOwner(c context.Context, user models.User, owner models.Owner) (res models.User, err error)
//	GetRatings(c context.Context, ownerId string) (res []interface{}, err error)
//	RegisterNewDriver(c context.Context, user models.User, driver models.Driver) (res models.User, err error)
//	GetDrivers(c context.Context, ownerId string) (res interface{}, err error)
//	GetOwnerStatusVerified(c context.Context, ownerId string) (res bool, err error)
//}
//
//type OwnerRepoImpl struct {
//	db *gorm.DB
//}
//
//func (a *OwnerRepoImpl) RegisterBusinessOwner(c context.Context, user models.User, owner models.Owner) (res models.User, err error) {
//	err = a.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
//		if err := tx.Create(&user).Error; err != nil {
//			return err
//		}
//
//		if err := tx.Create(&owner).Error; err != nil {
//			return err
//		}
//
//		return nil
//	})
//
//	if err != nil {
//		return res, err
//	}
//
//	return user, nil
//}
//
//func (a *OwnerRepoImpl) GetRatings(c context.Context, ownerId string) (res []interface{}, err error) {
//
//	row, err := a.db.WithContext(c).Table("reviews").
//		Select("Reviews.review, Reviews.star, Reviews.created_at AS review_date, Users.first_name AS reviewer_first_name, Users.last_name AS reviewer_last_name, Drivers.registration_number AS driver_registration_number").
//		Joins("JOIN Drivers ON Reviews.driver_id = driver_id").
//		Joins("JOIN BusinessOwners ON Drivers.owner_id = BusinessOwners.id").
//		Joins("JOIN Users ON Reviews.user_id = Users.id").
//		Where("BusinessOwners.id = ?", ownerId).
//		Rows()
//
//	defer row.Close()
//
//	if err != nil {
//		return res, err
//	}
//
//	type review struct {
//		Text               string
//		Star               int
//		Time               time.Time
//		UserFirstName      string
//		UserLastName       string
//		RegistrationNumber string
//	}
//
//	for row.Next() {
//		//_ = row.Scan(&user)
//		//users = append(users, user)
//	}
//
//	return res, nil
//}
//
//func (a *OwnerRepoImpl) RegisterNewDriver(c context.Context, user models.User, driver models.Driver) (res models.User, err error) {
//	err = a.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
//		if err := tx.Create(&user).Error; err != nil {
//			return err
//		}
//
//		if err := tx.Create(&driver).Error; err != nil {
//			return err
//		}
//
//		return nil
//	})
//
//	if err != nil {
//		return res, err
//	}
//
//	return user, nil
//}
//
//func (a *OwnerRepoImpl) GetDrivers(c context.Context, ownerId string) (res interface{}, err error) {
//	row, err := a.db.WithContext(c).Table("drivers").
//		Select("Users.first_name, Users.last_name, Drivers.registration_number").
//		Joins("JOIN Users ON Drivers.id = Users.id").
//		Where("Drivers.owner_id = ?", ownerId).
//		Rows()
//
//	defer row.Close()
//
//	type driver struct {
//		FirstName          string
//		LastName           string
//		RegistrationNumber string
//	}
//
//	var d driver
//	var users []driver
//	if err != nil {
//		return res, err
//	}
//
//	for row.Next() {
//		_ = row.Scan(&d.FirstName, &d.LastName, &d.RegistrationNumber)
//		users = append(users, d)
//	}
//
//	return users, nil
//}
//
//func (a *OwnerRepoImpl) GetOwnerStatusVerified(c context.Context, ownerId string) (res bool, err error) {
//	var owner models.Owner
//
//	if err := a.db.WithContext(c).First(&owner, "id = ?", ownerId).Error; err != nil {
//		return res, err
//	}
//
//	return owner.Verified, nil
//}
//
//func NewOwnerRepo(db *gorm.DB) OwnerRepo {
//	return &OwnerRepoImpl{
//		db: db,
//	}
//}
