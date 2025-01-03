package model

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
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	err = db.AutoMigrate(&DriverDetails{}, &Route{}, &Trip{})

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	return db
}
