package main

import (
	"rent/user-service/configs"
	"rent/user-service/src/controllers"
	"rent/user-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.User{})

	router := gin.Default()

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUser)

	router.Run("0.0.0.0:8081")
}
