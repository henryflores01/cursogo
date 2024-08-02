package funciones

import (
	"fmt"
)

// Recursion:
// Funcion que se llama asi misma n numero de veces
func Exponencia(valor int) {
	// Recuerda siempre poner al principio la salida de la recursion
	// para que no se llame infinitas veces
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}
