package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)

		w.Write([]byte(fmt.Sprintf("Hello world %d", number)))
	})

	http.ListenAndServe(":8080", nil)

}
