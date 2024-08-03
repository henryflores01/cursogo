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
	// "runtime"
	// "github.com/henryflores01/cursogo/variables"
	// "github.com/henryflores01/cursogo/ejercicios"
	// "github.com/henryflores01/cursogo/teclado"
	// "github.com/henryflores01/cursogo/iteraciones"
	// "github.com/henryflores01/cursogo/files"
	// "github.com/henryflores01/cursogo/funciones"
	// "github.com/henryflores01/cursogo/arreglos_slices"
	// "github.com/henryflores01/cursogo/mapas"
	// "github.com/henryflores01/cursogo/users"
	// "github.com/henryflores01/cursogo/ejer_interfaces"
	// "github.com/henryflores01/cursogo/modelos"
	"github.com/henryflores01/cursogo/defer_painic_recover"
)

// funcion principal
func main() {
	// imprimir en consola con un salto de linea
	// fmt.Println("Hola mundo")
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoATexto(1588)
	// fmt.Println(estado, texto)

	// IF - IF ELSE - ELSE
	// os := runtime.GOOS
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Esto no es windows, es: ", os)
	// } else if os == "darwin" {
	// 	fmt.Println("Esto no es windows, es: ", os)
	// } else {
	// 	fmt.Println("Esto es windows")
	// }

	// // Switch
	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Es linux")
	// case "darwin":
	// 	fmt.Println("Es darwin")
	// case "windows":
	// 	fmt.Println("Es windows")
	// default:
	// 	fmt.Printf("Es: %s\n", os)
	// }

	// numero, mensaje := ejercicios.ConvertirNumero("21")
	// fmt.Println(numero, mensaje)

	// captura de datos por teclado
	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.TablaDeMultiplicar())

	// files.GrabarTabla(ejercicios.TablaDeMultiplicar())
	// files.SumarTabla(ejercicios.TablaDeMultiplicar())
	// files.LeoArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencia(2)

	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestroSlice()
	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()
	// users.AltaUsuario()

	// Henry := new(modelos.Hombre)
	// Clari := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Henry)
	// ejer_interfaces.HumanosRespirando(Clari)

	// defer_painic_recover.VemosDefer()
	// defer_painic_recover.EjemploPanic()
	defer_painic_recover.EjemploRecover()

}
