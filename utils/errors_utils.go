package utils

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}

func WriteErrtoHttp(w http.ResponseWriter, appError *AppError) {
	jsonValue, _ := json.Marshal(appError)
	w.WriteHeader(appError.StatusCode)
	w.Write(jsonValue)
	return
}
