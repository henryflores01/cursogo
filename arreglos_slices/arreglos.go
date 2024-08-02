package arreglos_slices

import (
	"fmt"
)

func MuestroArreglos() {
	// Arreglos o vectores

	// 0 - 9
	var tabla [10]int

	// asignar valor
	tabla[0] = 33
	tabla[1] = 11

	// las posiciones que no tengan valor por defecto tienen 0
	fmt.Println(tabla)

	// crear y asignar
	var tabla2 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(tabla2)

	// asignar valor 2
	tabla3 := [3]string{"Henry", "Clari", "Jerry"}
	fmt.Println(tabla3)

	// recorrer array con for
	for i := 0; i < len(tabla3); i++ {
		fmt.Println(tabla3[i])
	}

	// matriz multidimensional
	var matriz [100][12]int

	matriz[20][10] = 20
	fmt.Println(matriz)
}
