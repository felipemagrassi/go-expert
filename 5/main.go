package main

import "fmt"

func main() {
	s := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(s)
	s["four"] = 4
	fmt.Println(s)
	delete(s, "four")
	fmt.Println(s)

	s1 := make(map[string]int)
	s2 := map[string]int{}
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)

	for key, value := range s {
		fmt.Println("key: ", key, "value: ", value)
	}

	for _, value := range s {
		fmt.Println("value: ", value)
	}
}
