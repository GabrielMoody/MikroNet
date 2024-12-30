package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DatabaseInit() *gorm.DB {

	dsn := fmt.Sprintf("host=localhost user=postgres password=mysecretpassword dbname=driver port=5432 sslmode=disable TimeZone=Asia/Singapore")

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
