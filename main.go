package main

import (
	"log"
	"startup-apps/auth"
	"startup-apps/handlers"
	"startup-apps/users"

	"github.com/gin-gonic/gin"
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
	userService := users.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handlers.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.POST("/users/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("users/avatar", userHandler.UploadAvatar)
	router.Run()
}