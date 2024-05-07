package models

import (
	"database/sql"
)

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Correo        string `json:"correo"`
	Contrasenia   string `json:"contrasenia"`
	FechaRegistro string `json:"fecha_registro"`
}

type UserService struct {
	DB *sql.DB
}

func (us UserService) CreateUser(u User) error {
	_, err := us.DB.Exec("INSERT INTO usuarios (nombre, correo, contrasenia, fecha_registro) VALUES (?,?,?,CURRENT_TIMESTAMP)", u.Nombre, u.Correo, u.Contrasenia)
	if err != nil {
		return err
	}
	return nil
}
