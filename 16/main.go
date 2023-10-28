package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("http://www.goole.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // schedule to close the body when the function returns similar to ensure in ruby

	res, err := io.ReadAll(req.Body)

	println(string(res))
}
