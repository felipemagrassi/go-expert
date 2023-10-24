package main

import (
	"go-expert/matematica"

	"github.com/google/uuid"
)

func main() {
	carro := matematica.Carro{Marca: "Fiat"}
	uuid := uuid.New()

	println(matematica.Soma(3, 2)) // capital letter, visible outside matematica package (public)
	// println(matematica.soma(3, 2)) // only visible inside matematica package (private)
	println(matematica.A) // capital letter, visible outside matematica package (public)
	// println(matematica.b) // only visible inside matematica package (private)
	println(carro.Marca)
	// println(carro.cor) // only visible inside matematica package (private)
	println(carro.Andar()) // capital letter, visible outside matematica package (public)
	println(uuid.String())
}
