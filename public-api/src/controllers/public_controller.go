package controllers

import (
	"fmt"
	"net/http"
	"rent/public-api/src/dtos"
	"rent/public-api/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllListings(ctx *gin.Context) {
	pageNumStr := ctx.DefaultQuery("page_num", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Create a map to hold the query parameters
	queryParams := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}

	body, _, err := services.ApiGet("http://listing-service:8080/listings", queryParams)
	if err != nil {
		// Handle error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, body)
}

func CreateUser(ctx *gin.Context) {
	var body dtos.CreateUserDto
	ctx.BindJSON(&body)

	responseBody, statusCode, err := services.ApiPost("http://user-service:8081/users", body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit data"})
		return
	}

	ctx.IndentedJSON(statusCode, responseBody)
}

func CreateListing(ctx *gin.Context) {
	var body dtos.CreateListingDto
	ctx.BindJSON(&body)

	user, userStatusCode, err := services.ApiGet(fmt.Sprintf("http://user-service:8081/users/%d", body.UserId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	if userStatusCode != http.StatusOK {
		ctx.JSON(userStatusCode, user)
		return
	}

	createListing, listingStatusCode, err := services.ApiPost("http://listing-service:8080/listings", body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit data"})
		return
	}

	ctx.IndentedJSON(listingStatusCode, createListing)
}
