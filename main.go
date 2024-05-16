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
	UserService := &models.UserService{
		DB: db,
	}

	UserController := &controllers.UserController{
		UserService: *UserService,
	}

	/*MiddleWareController := controllers.MiddleWareController{
		MiddleWareService: UserService,
	}*/

	VehicleService := &models.VehicleService{
		VehicleService: UserService,
		DB:             db,
	}

	VehicleController := &controllers.VehicleController{
		VehicleService: VehicleService,
	}

	DeviceService := &models.DeviceService{
		DB: db,
	}

	DeviceController := &controllers.DeviceController{
		DeviceService: DeviceService,
	}

	AlertService := &models.AlertService{
		AlertService: DeviceService,
		DB:           db,
	}

	AlertController := &controllers.AlertController{
		AlertService: AlertService,
	}
	/*c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
	})

	// Envuelve tu manejador HTTP con el controlador CORS
	handler := c.Handler(http.DefaultServeMux)*/
	handler := cors.Default().Handler(http.DefaultServeMux)
	http.HandleFunc("/createUser", UserController.CreateUser)
	http.HandleFunc("/loginUser", UserController.LoginUser)
	http.HandleFunc("/createVehicle", VehicleController.CreateVehicle)
	http.HandleFunc("/getvehicles", VehicleController.GetVehicles)
	http.HandleFunc("/getVehicle", VehicleController.GetVehicle)
	http.HandleFunc("/linkDevice", DeviceController.LinkDevice)
	http.HandleFunc("/createAlert", AlertController.CreateAlert)
	http.HandleFunc("/getAlerts", AlertController.GetAlerts)
	http.ListenAndServe("0.0.0.0:8080", handler)
}
