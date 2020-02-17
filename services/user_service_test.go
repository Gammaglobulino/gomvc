package services

import (
	"../models"
	"../services"
	"../utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDtoMock     usersDtoMock
	getUserFunction func(userId int64) (*models.User, *utils.AppError)
)

func init() {
	models.UserDto = &usersDtoMock{}
}

type usersDtoMock struct{}

func (m *usersDtoMock) GetUser(userId int64) (*models.User, *utils.AppError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*models.User, *utils.AppError) {
		return nil, &utils.AppError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := services.UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*models.User, *utils.AppError) {
		return &models.User{
			Id:        123,
			FirstName: "Andrea",
			LastName:  "Mazzanti",
			Email:     "mazzantia@hotmail.com",
		}, nil
	}
	user, err := services.UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
