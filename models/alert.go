package models

import (
	"database/sql"
	"fmt"
)

type Alert struct {
	ID            int     `json:"id_alerta"`
	IDDispositivo int     `json:"id_dispositivo"`
	Latitud       float64 `json:"latitud"`
	Longitud      float64 `json:"longitud"`
	Fecha_Alerta  string  `json:"fecha"`
	Estado        int     `json:"estado"`
}

type AlertService struct {
	AlertService *DeviceService
	DB           *sql.DB
}

func (as *AlertService) CreateAlert(id_dispositivo int, latitud float64, longitud float64) error {
	_, err := as.DB.Exec("INSERT INTO alertas (id_dispositivo, latitud, longitud) VALUES (?,?,?)", id_dispositivo, latitud, longitud)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (as *AlertService) ShowAlert(id_vehiculo int) (Alert, error) {
	id_dispositivo, err := as.AlertService.GetIDDevice(id_vehiculo)
	if err != nil {
		fmt.Println(err.Error())
		return Alert{}, err
	}

	var alert Alert
	row := as.DB.QueryRow("SELECT * FROM alertas WHERE id_dispositivo=?", id_dispositivo)
	err = row.Scan(&alert.ID, &alert.IDDispositivo, &alert.Latitud, &alert.Longitud, &alert.Fecha_Alerta, &alert.Estado)
	if err != nil {
		fmt.Println("Error al escanear:", err)
		return Alert{}, err
	}

	return alert, nil
}
