package interfaces

// Interfaces:
// modelos de comportamiento que podemos asociar
// a nuestras estructuras

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Nombre() string
	// EstaVivo() bool
}
