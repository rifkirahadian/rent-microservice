package controllers

import (
	"net/http"
	"rent/listing-service/configs"
	"rent/listing-service/src/dtos"
	"rent/listing-service/src/models"
	"rent/listing-service/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateListing(ctx *gin.Context) {
	var body dtos.CreateListingDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}

	if body.ListingType != "rent" && body.ListingType != "sale" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Listing type should be between rent/sale"})
		return
	}

	db := configs.InitDB()

	listing := models.Listing{
		UserId:      body.UserId,
		ListingType: body.ListingType,
		Price:       body.Price,
	}

	db.Create(&listing)

	response := dtos.ListingResponseDto{
		Result:  true,
		Listing: listing,
	}

	ctx.IndentedJSON(http.StatusCreated, response)
}

func GetAllListing(ctx *gin.Context) {
	db := configs.InitDB()
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

	var listings []models.Listing
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&listings)

	response := dtos.ListingsResponseDto{
		Result:   true,
		Listings: listings,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
