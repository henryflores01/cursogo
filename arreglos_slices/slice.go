package arreglos_slices

import (
	"fmt"
)

func MuestroSlice() {
	// definiendo slice: es un slice porque no tiene dimension fija
	// es un elemento dinamico
	var tabla []int = []int{2, 5, 4}
	fmt.Println(tabla)

	// crear un slice, tomando un arreglo como modelo
	// todos los slice apuntan a un arreglo mas grande
	var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	porcion := arreglo[3:]   // slice creado desde un vector desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]  // slice creado desde un vector desde la posicion 0 hasta la 5
	porcion3 := arreglo[4:7] // desde la posicion 4 hasta la 7

	fmt.Println(porcion, porcion2, porcion3)

}

func Capacidad() {
	// los slices manejan dos valores, largo y capacidad
	// capacidad: el largo del arreglo al que apunta
	// largo: el tamano del slice

	elementos := make([]int, 5, 20)
	fmt.Println("Largo: ", len(elementos))
	fmt.Println("Capacidad: ", cap(elementos))

	numeros := make([]int, 0, 0)

	// a medida que va creciendo se multiplica su capacidad
	for i := 0; i < 100; i++ {
		numeros = append(numeros, i)

		fmt.Println(numeros)
		fmt.Println("Largo: ", len(numeros))
		fmt.Println("Capacidad: ", cap(numeros))
	}
}
