package models

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Correo        string `json:"correo"`
	Contrasenia   string `json:"contrasenia"`
	FechaRegistro string `json:"fecha_registro"`
	Token         string `json:"token"`
}

type UserService struct {
	DB *sql.DB
}

func (us UserService) CreateUser(u User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Contrasenia), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Contrasenia = string(hashedPassword)

	_, err = us.DB.Exec("INSERT INTO usuarios (nombre, correo, contrasenia, fecha_registro) VALUES (?,?,?,CURRENT_TIMESTAMP)", u.Nombre, u.Correo, u.Contrasenia)
	if err != nil {
		return err
	}

	return nil
}

func (us UserService) Login(nombre, contrasenia string) (int, error) {
	var usuario User
	err := us.DB.QueryRow("SELECT id_usuario, nombre, correo, contrasenia, fecha_registro FROM usuarios WHERE nombre = ?", nombre).Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo, &usuario.Contrasenia, &usuario.FechaRegistro)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("usuario no encontrado")
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasenia), []byte(contrasenia))
	if err != nil {
		return 0, fmt.Errorf("contraseña incorrecta")
	}

	return usuario.ID, nil
}

func (us UserService) CrearSesion(ID int, token string) error {
	var count int
	err := us.DB.QueryRow("SELECT COUNT(*) FROM sesiones WHERE id_usuario = ?", ID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	_, err = us.DB.Exec("INSERT INTO sesiones (id_usuario, token) VALUES (?, ?)", ID, token)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) ConsultaID(token string) (int, error) {
	var id int
	err := us.DB.QueryRow("SELECT id_usuario FROM sesiones WHERE token = ?", token).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("token not found: %v", err)
		}
		return 0, err
	}
	return id, nil
}

func (us *UserService) CerrarSesion(token string) error {
	fmt.Println("Este es el token: " + token)
	_, err := us.DB.Exec("DELETE FROM sesiones WHERE token=?", token)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
