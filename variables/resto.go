package variables

import (
	"fmt"
	"strconv"
	"time"
)

// variable global del paquete
var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Henry"
	Estado = true
	Sueldo = 10000.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}
