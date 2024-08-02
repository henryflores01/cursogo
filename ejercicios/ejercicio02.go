package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func PedirNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingresa un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(numero * i)
	}
}
