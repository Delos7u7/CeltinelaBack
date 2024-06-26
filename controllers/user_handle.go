package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type ResponseJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseJsonNumber struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Telefono string `json:"telefono"`
}

type UserController struct {
	UserService models.UserService
}

type Claims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
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
	fmt.Println(userMode)
	w.Write(usersJson)
}

func (uc UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userMode models.User
	err := json.NewDecoder(r.Body).Decode(&userMode)
	if err != nil {
		fmt.Println(err.Error())
	}
	userID, err := uc.UserService.Login(userMode.Nombre, userMode.Contrasenia)
	if err != nil {
		fmt.Println(err.Error())
	}

	claims := Claims{
		UserID:         userID,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("asaksksja"))
	fmt.Println(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uc.UserService.CrearSesion(userID, tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": tokenString}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (uc UserController) CerrarSesion(w http.ResponseWriter, r *http.Request) {
	var userMode models.User

	err := json.NewDecoder(r.Body).Decode(&userMode)
	if err != nil {
		fmt.Println(err.Error())
	}

	token := userMode.Token

	err = uc.UserService.CerrarSesion(token)
	if err != nil {
		fmt.Println(err.Error())
	}

	responseCerrarSeion := ResponseJson{
		Code:    200,
		Message: "Sesion cerrada",
	}

	cerrarSesionJson, err := json.Marshal(responseCerrarSeion)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cerrarSesionJson)

}

func (uc UserController) Saludar(w http.ResponseWriter, r *http.Request) {
	responseCerrarSeion := ResponseJson{
		Code:    200,
		Message: "Sesion cerrada",
	}

	cerrarSesionJson, err := json.Marshal(responseCerrarSeion)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cerrarSesionJson)
}
