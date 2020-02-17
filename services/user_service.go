package services

import (
	"../models"
	"../utils"
)

type userService struct{}

var (
	UserService userService
)

func (s *userService) GetUser(userId int64) (*models.User, *utils.AppError) {
	return models.UserDto.GetUser(userId)
}
