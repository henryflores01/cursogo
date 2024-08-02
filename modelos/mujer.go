package modelos

type Mujer struct {
	Hombre  // herencia
	esMadre bool
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Pensar() {
	m.pensando = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Nombre() string {
	return m.nombre
}
