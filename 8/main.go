package main

import "fmt"

func main() {
	a := 10
	b := 20

	c := &a
	d := &b

	// types
	fmt.Printf("a: %T, c: %T\n", a, c)
	fmt.Printf("b: %T, d: %T\n", b, d)

	*c = 40

	fmt.Printf("a: %d, c: %d\n", &a, c)
	fmt.Printf("a: %d, c: %d\n", a, *c)
	fmt.Printf("b: %d, d: %d\n", &b, d)
	fmt.Printf("b: %d, d: %d\n", b, *d)

	fmt.Println(sum(a, b))
	fmt.Println(sumWithPointer(&a, &b))
	fmt.Println(sum(a, b))
}

func sum(a, b int) int {
	fmt.Printf("a: %d, b: %d\n", &a, &b)
	return a + b
}

func sumWithPointer(a, b *int) int {
	fmt.Printf("a: %d, b: %d\n", a, b)
	*a = 0
	return *a + *b
}
