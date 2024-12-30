package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DatabaseInit() *gorm.DB {

	dsn := fmt.Sprintf("root:123@tcp(127.0.0.1:3307)/mikronet?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error while connecting database"))
	}

	log.Print("Connection Succeed")

	err = db.AutoMigrate(&OwnerDetails{}, &BlockedAccount{}, &GovDetails{}, &Admin{}, &OwnerDetails{})

	if err != nil {
		panic(fmt.Errorf("error while migrating database"))
	}

	return db
}
