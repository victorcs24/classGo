package main

import "fmt"

var tabla [10]int

//var matriz [5][7]int
func main() {
	fmt.Println("Arreglos ---- vectores")

	tabla[0] = 1
	tabla[5] = 6
	fmt.Println(tabla)
	fmt.Println("Arreglos --7-- vectores tabla2")

	tabla2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(tabla2)
	fmt.Println("recorrido arreglos")
	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	fmt.Println("funcion main")
	makeslice()
}

func makeslice() {
	elemento := make([]int, 0, 0)
	for i := 0; i < 129; i++ {
		elemento = append(elemento, i)
	}
	fmt.Printf("slice tamaño %d capacidad %d", len(elemento), cap(elemento))
}
