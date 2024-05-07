package main

import (
	"CeltinelaBack/controllers"
	"CeltinelaBack/models"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var db *sql.DB

func ConexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "123456"
	Nombre := "celtinela"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		fmt.Println(err.Error())
	}
	return conexion
}

func main() {
	db = ConexionBD()
	defer db.Close()
	UserService := models.UserService{
		DB: db,
	}

	UserController := controllers.UserController{
		UserService: UserService,
	}
	handler := cors.Default().Handler(http.DefaultServeMux)
	http.HandleFunc("/createUser", UserController.CreateUser)
	http.ListenAndServe(":8080", handler)
}
