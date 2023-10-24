package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

type Carro struct {
	Marca string
	cor   string
}

func (c Carro) Andar() int {
	println("Andando...")

	return 1
}

var A int = 10
var b int = 20
