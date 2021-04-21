package main

import "fmt"

//var numero int
//var texto string
//var status bool

//float32 75334395
var numero1, numero2 int

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("\n variable funcion")
	fmt.Println(calculo(2, 4))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println("\n variable funcion redefinir resta ")
	fmt.Println(calculo(10, 3))
	fmt.Println("\n ")
	operaciones()
	fmt.Println("\n funcion closures ")
	tablaDel := 2
	miTabla := Tabla(tablaDel)

	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}
func operaciones() {
	resultado := func() int {
		var a int = 5
		var b int = 6
		return a * b
	}
	fmt.Printf("\n resultado operaciones %d", resultado())
}

func Tabla(valor int) func() int {
	numT := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numT * secuencia
	}
}

func main_funciones() {
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
