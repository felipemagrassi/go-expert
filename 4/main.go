package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s), cap(s), s)

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:]), cap(s[:]), s[:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[1:]), cap(s[1:]), s[1:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[2:]), cap(s[2:]), s[2:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:1]), cap(s[:1]), s[:1])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[2:2]), cap(s[2:2]), s[2:2])

	s = append(s, 10)

	fmt.Println("----------------------")

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s), cap(s), s)

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:]), cap(s[:]), s[:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[1:]), cap(s[1:]), s[1:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[2:]), cap(s[2:]), s[2:])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:1]), cap(s[:1]), s[:1])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len: %d, cap: %d, value: %v\n", len(s[2:2]), cap(s[2:2]), s[2:2])
}
