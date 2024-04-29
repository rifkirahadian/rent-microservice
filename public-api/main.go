package main

import (
	"rent/public-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/public-api/listings", controllers.GetAllListings)
	router.POST("/public-api/users", controllers.CreateUser)
	router.POST("/public-api/listings", controllers.CreateListing)

	router.Run("0.0.0.0:8082")
}
