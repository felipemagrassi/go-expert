package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compare[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.0, "b": 2.0, "c": 3.0}
	m3 := map[string]MyNumber{"a": 1.0, "b": 2.0, "c": 3.0}

	println(SomaInt(m))
	println(SomaFloat(m2))
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compare(1, 1))
}
