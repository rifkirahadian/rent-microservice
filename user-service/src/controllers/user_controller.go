package controllers

import (
	"net/http"
	"rent/user-service/configs"
	"rent/user-service/src/dtos"
	"rent/user-service/src/models"
	"rent/user-service/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var body dtos.CreateUserDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}
	db := configs.InitDB()

	user := models.User{
		Name: body.Name,
	}

	db.Create(&user)

	response := dtos.UserResponseDto{
		Result: true,
		User:   user,
	}

	ctx.IndentedJSON(http.StatusCreated, response)
}

func GetAllUsers(ctx *gin.Context) {
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

	var users []models.User
	db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users)

	response := dtos.UsersResponseDto{
		Result: true,
		Users:  users,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	db := configs.InitDB()
	var user models.User

	db.First(&user, id)

	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	response := dtos.UserResponseDto{
		Result: true,
		User:   user,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}
