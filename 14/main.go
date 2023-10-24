package main

func main() {
	a := 9
	b := 2

	if a == 1 {
		println(a)
	} else {
		println(b)
	}

	if (a == 1 && b == 2) || a == 2 {
		println(b)
	} else {
		println("hello")
	}

	switch a {
	case 1:
		println(a)
	case 3:
		println(123)
	default:
		println(b)
	}

}
