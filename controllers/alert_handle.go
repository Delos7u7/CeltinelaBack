package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type AlertController struct {
	AlertService *models.AlertService
}

func (ac AlertController) CreateAlert(w http.ResponseWriter, r *http.Request) {
	id_dispositivostr := r.URL.Query().Get("id_dispositivo")
	latitudstr := r.URL.Query().Get("latitud")
	longitudstr := r.URL.Query().Get("longitud")
	id_dispositivo, err := strconv.Atoi(id_dispositivostr)
	if err != nil {
		fmt.Println(err.Error())
	}
	latitud, err := strconv.ParseFloat(latitudstr, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	longitud, err := strconv.ParseFloat(longitudstr, 64)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ac.AlertService.CreateAlert(id_dispositivo, latitud, longitud)
	if err != nil {
		http.Error(w, "Error al crear la alerta", http.StatusBadRequest)
		return
	}
	responseCreateAlert := ResponseJson{
		Code:    200,
		Message: "Alerta Creada",
	}

	jsonResponseCreateAlert, err := json.Marshal(responseCreateAlert)

	if err != nil {
		http.Error(w, "Error el json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonResponseCreateAlert)

}

func (ac AlertController) GetAlerts(w http.ResponseWriter, r *http.Request) {
	id_vehiculostr := r.URL.Query().Get("id_vehiculo")
	id_vehiculo, err := strconv.Atoi(id_vehiculostr)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error al obtener el ID del vehículo", http.StatusBadRequest)
		return
	}

	alert, err := ac.AlertService.ShowAlerts(id_vehiculo)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error al obtener la alerta", http.StatusInternalServerError)
		return
	}

	jsonAlert, err := json.Marshal(alert)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error al convertir la alerta a JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonAlert)
}

func (ac AlertController) ChangeAlertState(w http.ResponseWriter, r *http.Request) {
	id_vehiculostr := r.URL.Query().Get("id_vehiculo")
	id_vehiculo, err := strconv.Atoi(id_vehiculostr)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error al obtener el ID del vehículo", http.StatusBadRequest)
		return
	}

	err = ac.AlertService.ChangeAlertState(id_vehiculo)
	if err != nil {
		fmt.Println(err.Error())
	}

	responseChangeAlertState := ResponseJson{
		Code:    200,
		Message: "Estado cambiado",
	}

	jsonResponseChangeAlertState, err := json.Marshal(responseChangeAlertState)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(jsonResponseChangeAlertState)
}
