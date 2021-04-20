package main

import "fmt"

//var numero int
//var texto string
//var status bool

//float32 75334395
var numero1, numero2 int

func main() {
	fmt.Printf("numero %d", doblenum(4))

	numx, result := dobleparam(3)
	fmt.Println("\n Doble parametro")
	fmt.Println(numx)
	fmt.Println(result)
	fmt.Println("\n Metodo calcular")
	calcular(2, 4, 6, 7)
	fmt.Println("\n Metodo calcular")
	calcular(7, 4, 5)

}

func doblenum(numero int) (z int) {
	z = numero * 2
	return z
}

func dobleparam(numero int) (int, bool) {
	if numero == 2 {
		return (numero * 2), true
	} else {
		return (numero * 3), false
	}
}

func calcular(numero ...int) int {
	total := 0

	//for _, num := range numero {
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n item: %d total: %d", item, total)
	}
	return total
}

func calcular2(numero ...int) int {
}

func main_X2() {
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
