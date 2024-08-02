package interfaces

// Interfaces:
// modelos de comportamiento que podemos asociar
// a nuestras estructuras

type Animal interface {
	Respirar()
	Comer()
	EsCarnivoro() bool
	EstaVivo() bool
}
