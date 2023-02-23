package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(w http.ResponseWriter, status bool, message string, data any) {
	jsonMap := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	res, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}

func BuildErrorResponse(w http.ResponseWriter, message string, errors string, data any) {
	jsonMap := Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
	res, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, string(res))
}
