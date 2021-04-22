package main

import (
	"fmt"
	"time"
)

type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func main() {
	User := new(usuario)
	User.Id = 2
	User.Nombre = "Chano"
	fmt.Println(User)
}
