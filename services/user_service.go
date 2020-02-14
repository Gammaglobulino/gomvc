package services

import (
	"../models"
	"../utils"
)

func GetUser(userId int64) (*models.User, *utils.AppError) {
	return models.GetUser(userId)
}
