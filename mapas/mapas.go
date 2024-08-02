package mapas

import (
	"fmt"
)

func MostrarMapas() {
	// paises := make(map[string]string)
	// // fmt.Println(paises)

	// paises["Mexico"] = "CDMX"
	// paises["Argentina"] = "Buenos Aires"

	// fmt.Println(paises)
	// fmt.Println(paises["Mexico"])

	champions := map[string]int{
		"Real Madrid": 15,
		"Milan":       7,
		"Barcelona":   5,
	}

	fmt.Println(champions)
	fmt.Println(champions["Real Madrid"])

	for clave, valor := range champions {
		fmt.Println(clave, valor)
	}

	// eliminar elemento
	delete(champions, "Barcelona")
	fmt.Println(champions)

	// saber si existe un valor
	championsTotal, exite := champions["Juventus"]
	fmt.Println(championsTotal, exite)

	championsTotal, exite = champions["Real Madrid"]
	fmt.Println(championsTotal, exite)
}
