package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserController struct {
	UserService models.UserService
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userMode models.User
	err := json.NewDecoder(r.Body).Decode(&userMode)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = uc.UserService.CreateUser(userMode)
	if err != nil {
		fmt.Println(err.Error())
	}

	responseCreateUser := ResponseJson{
		Code:    200,
		Message: "Usuario Creado",
	}

	usersJson, err := json.Marshal(responseCreateUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(usersJson)
}
