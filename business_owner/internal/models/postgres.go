package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func DatabaseInit() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASS"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	//err = db.AutoMigrate(&User{}, &ResetPassword{})

	//if err != nil {
	//	panic(fmt.Errorf("error while migrating database"))
	//}

	return db
}
