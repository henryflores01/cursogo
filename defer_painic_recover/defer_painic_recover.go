package defer_painic_recover

import (
	"fmt"
	"log"
)

// DEFER:
// Herramienta que nos da GO que nos permite indicar cual va a ser
// la ultima instruccion en ejectuarse cuando salga de la funcion

func VemosDefer() {
	// defer archivo.Close() // eso haria que fuera la ultima cosa en ejecutarse

	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el mensaje final") // se ejectua al final

	fmt.Println("Este es el segundo mensaje")

}

// Aborta el programa con un mensaje en consola
// cuando hay alguna falla
func EjemploPanic() {
	a := 1
	if a == 1 {
		panic("Error, se encontro el valor 1")
	}
}

// Si se ejecuta un panic, permite recuperar el programa
// para despues manejar el error de otra manera
func EjemploRecover() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero un panic: \n %v\n", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Error, se encontro el valor 1")
	}
}
