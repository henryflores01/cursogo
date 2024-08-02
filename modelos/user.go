package modelos

import (
	"time"
)

type User struct {
	id        int
	nombre    string
	apellido  string
	createdAt time.Time
	status    bool
}

// solo las funciones pertenecientes a la estructura van en este
// archivo llamado users.go
// no olvidar la referencia a la estructura "(u User)"
func (u *User) AddUser(id int, nombre string, apellido string, createdAt time.Time, status bool) {
	u.id = id
	u.nombre = nombre
	u.apellido = apellido
	u.createdAt = createdAt
	u.status = status
}
