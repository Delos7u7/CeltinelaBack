package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type VehicleController struct {
	VehicleService *models.VehicleService
}

func (vc *VehicleController) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehiclesMode models.Vehiculo

	err := json.NewDecoder(r.Body).Decode(&vehiclesMode)
	fmt.Println(vehiclesMode)
	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON", http.StatusBadRequest)
		//fmt.Println(err.Error())
		return
	}

	token := vehiclesMode.Token
	err = vc.VehicleService.CreateVehicle(vehiclesMode, token)
	if err != nil {
		http.Error(w, "Error al crear el vehículo: "+err.Error(), http.StatusInternalServerError)
		//fmt.Println(err.Error())
		return
	}

	responseCreateVehicle := ResponseJson{
		Code:    200,
		Message: "Vehiculo Creado",
	}

	vehicleJson, err := json.Marshal(responseCreateVehicle)
	if err != nil {
		http.Error(w, "Error al serializar la respuesta JSON", http.StatusInternalServerError)
		//fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(vehicleJson)
}

func (vc *VehicleController) GetVehicles(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	arrayVehiculos, err := vc.VehicleService.SelectVehicles(token)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	jsonArrayVehiculos, err := json.Marshal(arrayVehiculos)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonArrayVehiculos)
}

func (vc *VehicleController) GetVehicle(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	idVehiculoStr := r.URL.Query().Get("idVehiculo")
	idVehiculo, err := strconv.Atoi(idVehiculoStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	vehiculo, err := vc.VehicleService.SelectVehicle(token, idVehiculo)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	jsonVehiculo, err := json.Marshal(vehiculo)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonVehiculo)
}
