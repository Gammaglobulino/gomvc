package controllers

import (
	"../services"
	"../utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("user_id")
	userId, errno := strconv.ParseInt(userIdParam, 10, 64)
	if errno != nil {
		userErr := &utils.AppError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.WriteErrtoHttp(w, userErr)
		return
	}
	user, err := services.UserService.GetUser(userId)
	if err != nil {
		utils.WriteErrtoHttp(w, err)
		return
	}
	jsonm, _ := json.Marshal(user)
	w.Write(jsonm)

}
