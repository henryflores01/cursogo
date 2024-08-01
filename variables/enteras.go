// El package es el nombre de la carpeta
package variables

import "fmt"

func MostrarEnteros() {
	// todas las variables creadas se deben usar o marcara error
	// las variables al crearlas no se definen como nulas, si no como el minimo de dato que pueden almacenar
	// como se declara una variable?

	// declarativa:
	// var name type
	var intcomun int

	// asignativa:
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println(intcomun)
	fmt.Println(intde32)
	fmt.Println(intde64)

}
