package models

import (
	"../utils"
	"fmt"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{
			Id:        123,
			FirstName: "Andrea",
			LastName:  "Mazzanti",
			Email:     "mazzantia@hotmail.com",
		},
	}
	UserDto usersServiceInterface
)

func init() {
	UserDto = &userDto{}
}

type usersServiceInterface interface {
	GetUser(int642 int64) (*User, *utils.AppError)
}
type userDto struct{}

func (u *userDto) GetUser(userId int64) (*User, *utils.AppError) {
	log.Println("We are accessing the Database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
