package main

import (
	"fmt"

	"github.com/felipemagrassi/go-expert/api/configs"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(configs)
}
