package models

import (
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DatabaseInit() *gorm.DB {
	v := helper.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", v.GetString("MYSQL_USERNAME"), v.GetString("MYSQL_PASS"), v.GetString("MYSQL_HOST"), v.GetString("MYSQL_PORT"), v.GetString("MYSQL_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	err = db.AutoMigrate(&User{})

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	return db
}

func DatabaseTestInit() *gorm.DB {
	v := helper.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", v.GetString("MYSQL_USERNAME"), v.GetString("MYSQL_PASS"), v.GetString("MYSQL_HOST"), v.GetString("MYSQL_PORT"), v.GetString("MYSQL_DATABASE_TEST"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	err = db.AutoMigrate(&User{})
	db.Exec("DELETE FROM users")

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	return db
}
