package models

import (
	"fmt"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DatabaseInit() *gorm.DB {
	v := helper.LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore", v.GetString("POSTGRES_HOST"), v.GetString("POSTGRES_USERNAME"), v.GetString("POSTGRES_PASS"), v.GetString("POSTGRES_DATABASE"), v.GetString("POSTGRES_PORT"))

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
