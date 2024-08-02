package iteraciones

import (
	"fmt"
)

func Iterar() {
	// FOR:

	// for {
	// 	fmt.Println("Hola")
	// 	// break // romper bucle
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 100; i += 2 {
	// 	fmt.Println(i)
	// }

	// 	for i := 100; i > 0; i -= 5 {
	// 		fmt.Println(i)
	// 	}

	for i := 10; i > 1; i-- {
		if i == 6 {
			// break // romper bucle
			continue // saltar iteracion
		}

		fmt.Println(i)
	}
}
