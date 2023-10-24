package main

func main() {
	// simple for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"cero", "uno", "dos", "tres"}

	// for loop with range
	for k, v := range numeros {
		println(k, v)
	}

	// while loop

	j := 0
	for j < 10 {
		println(j)
		j++
	}

	// infinite loop
	for {
		println("Hola")
		break
	}

}
