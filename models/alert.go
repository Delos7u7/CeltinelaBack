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

func (as *AlertService) ShowAlerts(id_vehiculo int) ([]Alert, error) {
	id_dispositivo, err := as.AlertService.GetIDDevice(id_vehiculo)
	if err != nil {
		fmt.Println(err.Error())
	}

	var AlertArray []Alert
	rows, err := as.DB.Query("SELECT * FROM alertas WHERE id_dispositivo=? AND estado='0' ORDER BY fecha_hora DESC", id_dispositivo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Alert Alert
		if err := rows.Scan(&Alert.ID, &Alert.IDDispositivo, &Alert.Latitud, &Alert.Longitud, &Alert.Fecha_Alerta, &Alert.Estado); err != nil {
			fmt.Println("Error al escanear:", err)
			return nil, err
		}
		AlertArray = append(AlertArray, Alert)
	}
	return AlertArray, nil
}

func (as *AlertService) ChangeAlertState(id_vehiculo int) error {
	id_dispositivo, err := as.AlertService.GetIDDevice(id_vehiculo)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = as.DB.Exec("UPDATE alertas SET estado = '1' WHERE id_dispositivo = ?", id_dispositivo)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
