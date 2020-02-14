package models

import (
	"../utils"
	"fmt"
	"net/http"
)

var users = map[int64]*User{
	123: &User{
		Id:        123,
		FirstName: "Andrea",
		LastName:  "Mazzanti",
		Email:     "mazzantia@hotmail.com",
	},
}

func GetUser(userId int64) (*User, *utils.AppError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
