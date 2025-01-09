package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func DatabaseInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database %v", err.Error()))
	}

	err = db.AutoMigrate(&User{}, &ResetPassword{})

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	log.Print("Connection Succeed")

	return db
}
