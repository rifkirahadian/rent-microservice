package dtos

import "rent/user-service/src/models"

type CreateUserDto struct {
	Name string `json:"name" binding:"required"`
}

type UserResponseDto struct {
	Result bool        `json:"result"`
	User   models.User `json:"user"`
}

type UsersResponseDto struct {
	Result bool          `json:"result"`
	Users  []models.User `json:"users"`
}
