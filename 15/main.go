package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// write file
	f, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	size, err := f.WriteString("Hello World\n")

	if err != nil {
		panic(err)
	}

	size, err = f.Write([]byte("Hello World\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", size)

	f.Close()

	//read file

	file, err := os.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file)) // string converts byte array to string

	// read file with buffer

	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("test.txt")

	if err != nil {
		panic(err)
	}

}
