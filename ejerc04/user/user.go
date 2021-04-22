package user

import (
	"time"
)

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}
