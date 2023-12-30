package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Hello, world! You are the visitor number %d\n", number)))
	})

	http.ListenAndServe(":8080", nil)
}
