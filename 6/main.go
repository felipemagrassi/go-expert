package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(-1, 2))
	fmt.Println(infiniteSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	by2 := multiplyBy(2)
	by10 := multiplyBy(10)

	fmt.Printf("%T\n", by2)

	fmt.Println(by2(2))
	fmt.Println(by10(2))

	func() { fmt.Println("Anonymous function") }()
}

// Functions
func sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("a or b is less than 0")
	}
	return a + b, nil
}

// Functions with variadic parameters
func infiniteSum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// Closure

func multiplyBy(power int) func(int) int {
	return func(num int) int {
		return num * power
	}
}
