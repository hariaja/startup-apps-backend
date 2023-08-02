package main

import (
	"log"
	"startup-apps/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Connect to database mysql
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// If connection error
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	user := users.User{
		Name: "Test Simpan",
	}

	userRepository.Store(user)
}