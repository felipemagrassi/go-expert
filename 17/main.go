package main

func main() {
	defer println("1")
	println("2")
	println("3")
	println("4")

	// Output: 2 3 4 1
}
