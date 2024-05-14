package models

import (
	"database/sql"
	"fmt"
)

type Vehiculo struct {
	ID          int    `json:"id_vehiculo"`
	IDUsuario   int    `json:"id_usuario"`
	Alias       string `json:"alias_vehiculo"`
	Tipo        string `json:"tipo_vehiculo"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Año         int    `json:"año"`
	Color       string `json:"color"`
	Placa       string `json:"placa"`
	NumSerieVIN string `json:"num_serie_vin"`
	Token       string `json:"token"`
}

type VehicleService struct {
	VehicleService *UserService
	DB             *sql.DB
}

func (vs *VehicleService) CreateVehicle(v Vehiculo, token string) error {
	fmt.Println("1.5 ", token)
	id, err := vs.VehicleService.ConsultaID(token)
	if err != nil {
		//fmt.Println(err.Error())
		return err
	}
	_, err = vs.DB.Exec("INSERT INTO Vehículos (id_usuario, alias_vehiculo, tipo_vehiculo, marca, modelo, año, color, placa, num_serie_vin) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", id, v.Alias, v.Tipo, v.Marca, v.Modelo, v.Año, v.Color, v.Placa, v.NumSerieVIN)
	if err != nil {
		//fmt.Println(err.Error())
		return err
	}
	return nil
}

func (vs *VehicleService) SelectVehicles(token string) ([]Vehiculo, error) {
	id, err := vs.VehicleService.ConsultaID(token)
	if err != nil {
		return nil, err
	}
	//var Vehiculo Vehiculo
	var VehiculoArray []Vehiculo
	rows, err := vs.DB.Query("SELECT * FROM Vehículos WHERE id_usuario=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Vehiculo Vehiculo
		if err := rows.Scan(&Vehiculo.ID, &Vehiculo.IDUsuario, &Vehiculo.Alias, &Vehiculo.Tipo, &Vehiculo.Marca, &Vehiculo.Modelo, &Vehiculo.Año, &Vehiculo.Color, &Vehiculo.Placa, &Vehiculo.NumSerieVIN); err != nil {
			return nil, err
		}
		VehiculoArray = append(VehiculoArray, Vehiculo)
	}
	return VehiculoArray, nil
}
