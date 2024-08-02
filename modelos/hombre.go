package modelos

type Hombre struct {
	nombre     string
	edad       int
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}

func (h *Hombre) Pensar() {
	h.pensando = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Nombre() string {
	return h.nombre
}
