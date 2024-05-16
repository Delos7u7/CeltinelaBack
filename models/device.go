package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type Device struct {
	ID         string `json:"id_dispositivo"`
	IDVehiculo string `json:"id_vehiculo"`
}

type DeviceService struct {
	DB *sql.DB
}

func (dv DeviceService) LinkDevice(d Device) error {
	IDinr, err := strconv.Atoi(d.ID)
	if err != nil {
		return err
	}
	IdVehiculo, err := strconv.Atoi(d.IDVehiculo)
	if err != nil {
		return err
	}
	fmt.Println("Este es el id:", IDinr)
	fmt.Println("Este es el idV:", IdVehiculo)
	_, err = dv.DB.Exec("UPDATE dispositivos SET id_vehiculo = ? WHERE id_dispositivo = ?", IdVehiculo, IDinr)
	if err != nil {
		return err
	}
	return nil
}

func (dv DeviceService) GetIDDevice(id_vehiculo int) (int, error) {
	var id int
	err := dv.DB.QueryRow("SELECT id_dispositivo FROM dispositivos WHERE id_vehiculo=?", id_vehiculo).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("id not found: %v", err)
		}
		return 0, err
	}
	return id, nil
}
