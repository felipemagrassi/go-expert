package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64

func main() {
	m := &sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()

		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Hello World! %d\n", number)))
	})
	http.ListenAndServe(":8080", nil)
}
