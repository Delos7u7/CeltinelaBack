package controllers

import (
	"CeltinelaBack/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeviceController struct {
	DeviceService *models.DeviceService
}

func (dc DeviceController) LinkDevice(w http.ResponseWriter, r *http.Request) {
	var deviceMode models.Device
	err := json.NewDecoder(r.Body).Decode(&deviceMode)
	fmt.Println(deviceMode)
	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON", http.StatusBadRequest)
		//fmt.Println(err.Error())
		return
	}
	//idDevice := deviceMode.ID
	//idVehiculo := deviceMode.IDVehiculo
	fmt.Println(deviceMode)
	err = dc.DeviceService.LinkDevice(deviceMode)

	if err != nil {
		//http.Error(w, "Error al linkear el dispositivo: "+err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	responseLinkDevice := ResponseJson{
		Code:    200,
		Message: "Dispositivo Enlazado",
	}

	deviceJson, err := json.Marshal(responseLinkDevice)

	if err != nil {
		//http.Error(w, "Error al serializar la respuesta JSON", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(deviceJson)
}
