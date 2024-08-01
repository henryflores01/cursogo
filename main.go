// Por convencion el primer archivo siempre es el main.go

// todo comienza por package, por convencion
// en este caso como el archivo se llama main y esta en la carpeta
// principal siempre va a ser "package main"
// Puedes crear multiples paquetes e importarlos
// El package es el nombre de la carpeta
package main

// importamos el paquete fmt, es un paquete integrado
// import "fmt"

// para importar varios paquetes usamos un parentesis
import (
	// "fmt"
	"fmt"

	"github.com/henryflores01/cursogo/variables"
)

// funcion principal
func main() {
	// imprimir en consola con un salto de linea
	// fmt.Println("Hola mundo")
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado, texto)

}
