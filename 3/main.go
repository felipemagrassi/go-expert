package main

import "fmt"

func main() {
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr[len(arr)-1])

	fmt.Println("print using for loop")
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}

	fmt.Println("print using range")
	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
		fmt.Printf("index type: %T\n", index)
	}
}
