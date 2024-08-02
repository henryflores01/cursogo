package ejer_interfaces

import (
	"fmt"

	"github.com/henryflores01/cursogo/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy %s y estoy respirando\n", hu.Nombre())
}
