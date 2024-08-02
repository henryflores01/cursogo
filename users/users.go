package users

import (
	"fmt"
	"time"

	"github.com/henryflores01/cursogo/modelos"
)

func AltaUsuario() {
	// instanciar un objeto
	u := new(modelos.User)
	u.AddUser(1, "Henry", "Flores", time.Now().UTC(), true)
	fmt.Println(u)
}
