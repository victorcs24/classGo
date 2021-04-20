package main

import (
	"fmt"
)

//var numero int
//var texto string
//var status bool

//float32 75334395
var numero1, numero2 int

func main() {
	//fmt.Println("Hola Mundo")
	//var numero2 int
	//solo se hace una vez
	//numero3 := 4
	fmt.Println("Ingrese numero1: ")
	//fmt.Scanf("%d", &numero1)   //linux
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero2: ")
	fmt.Scanln(&numero2)
	if numero1 > numero2 {
		fmt.Printf("Suma num1 mayor %d", (numero1 + numero2))
	} else {
		fmt.Printf("Suma num1 menor [%d],[%d]", numero1, numero2)
	}

}
