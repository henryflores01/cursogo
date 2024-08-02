package files

import (
	"bufio"
	"fmt"

	// "io/ioutil"
	"os"
)

var filename string = "tabla.txt"

func GrabarTabla(texto string) {
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)

	archivo.Close()
}

func SumarTabla(texto string) {
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar el contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir el archivo: " + err.Error())
		return false
	}

	arch.Close()

	return true
}

// func LeoArchivo() {
// 	arch, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error al leer el archivo: " + err.Error())
// 		return
// 	}

// 	fmt.Println(string(arch))
// }

// func LeoArchivo() {
// 	arch, err := os.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error al leer el archivo: " + err.Error())
// 		return
// 	}

// 	fmt.Println(string(arch))
// }

func LeoArchivo() {
	arch, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(registro)
	}
}
