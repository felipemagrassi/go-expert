package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	resp, err := c.Post("http://google.com", "application/json", bytes.NewBuffer([]byte(`{"hello": "world"}`)))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, make([]byte, 1024))
}
