package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	// dsn := "host=localhost user=postgres password=123 dbname=mikronet port=5432 sslmode=disable TimeZone=Asia/Singapore"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	err = db.AutoMigrate(&DriverLocationLogs{}, &DriverLocation{})

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	return db
}
