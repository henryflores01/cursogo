package ejercicios

import (
	"strconv"
)

func ConvertirNumero(valor string) (int, string) {
	var texto string
	numero, _ := strconv.Atoi(valor)

	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}

	return numero, texto
}
