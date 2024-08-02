package funciones

import (
	"fmt"
)

// Funciones Anonimas: No solo se crean funciones de la manera
// tradicional, si no que se pueden crear como variables,
// y pasar como parametros, estas son las funciones anonimas

func Calculos() {
	var numero3 int = 32
	var numero4 int = 243

	// funcion  anonima asignada a una variable
	// pueden acceder a variables de su scope padre
	calculo := func(num1, num2 int) int {
		return num1 + num2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	// no se puede redefinir la estructura de una funcion anonima
	// solo su funcionamiento
	calculo = func(num1, num2 int) int {
		return num1 * num2
	}

	fmt.Println(calculo(10, 25))
}
