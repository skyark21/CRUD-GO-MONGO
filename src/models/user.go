package models

import "time"

//Estructura de Datos de Usuario
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
}

//Lista de Usuarios
type Users []*User
