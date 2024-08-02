package interfaces

// Interfaces:
// modelos de comportamiento que podemos asociar
// a nuestras estructuras

type Vegetal interface {
	Clasificacion() string
	EstaVivo() bool
}
