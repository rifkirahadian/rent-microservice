package main

import (
	"rent/listing-service/configs"
	"rent/listing-service/src/controllers"
	"rent/listing-service/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	db.AutoMigrate(models.Listing{})

	router := gin.Default()

	router.POST("/listings", controllers.CreateListing)
	router.GET("/listings", controllers.GetAllListing)

	router.Run("0.0.0.0:8080")
}
