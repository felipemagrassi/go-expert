package main

import "fmt"

// Using empty interface to create a function that can take any type

func main() {
	var x interface{} = 7
	var y interface{} = "abc"

	println(x)
	println(y)

	value, err := x.(int)
	println(value, err) // type assertion

	fmt.Println(x)
	fmt.Println(y)

	printType(x)
	printType(y)
}

func printType(t interface{}) {
	fmt.Printf("%T\n", t)
}
