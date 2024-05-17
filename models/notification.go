package models

import (
	"database/sql"
	"fmt"
)

type Notification struct {
	Id          int    `json:"id"`
	Alias       string `json:"alias"`
	Tipo        string `json:"tipo_vehiculo"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Año         int    `json:"año"`
	Color       string `json:"color"`
	Placa       string `json:"placa"`
	NumSerieVIN string `json:"num_serie_vin"`
}

type NotificationService struct {
	NotificationService *UserService
	DB                  *sql.DB
}

func (ns NotificationService) Notification(token string) ([]Notification, error) {
	//fmt.Println("Este es el token: " + token)
	id, err := ns.NotificationService.ConsultaID(token)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//fmt.Println("Este es el id: ", id)
	rows, err := ns.DB.Query(`
        SELECT v.id_vehiculo, v.alias_vehiculo, v.tipo_vehiculo, v.marca, v.modelo, v.año, v.color, v.placa, v.num_serie_vin
        FROM Usuarios u
        JOIN Vehículos v ON u.id_usuario = v.id_usuario
        JOIN Dispositivos d ON v.id_vehiculo = d.id_vehiculo
        JOIN Alertas a ON d.id_dispositivo = a.id_dispositivo
        WHERE a.estado = '0' AND u.id_usuario = ?
    `, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var notificaciones []Notification
	for rows.Next() {
		var notificacion Notification
		err = rows.Scan(&notificacion.Id, &notificacion.Alias, &notificacion.Tipo, &notificacion.Marca, &notificacion.Modelo, &notificacion.Año, &notificacion.Color, &notificacion.Placa, &notificacion.NumSerieVIN)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		notificaciones = append(notificaciones, notificacion)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(notificaciones)
	return notificaciones, nil
}
