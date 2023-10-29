package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("request started")
	defer log.Println("request completed")

	select {
	case <-time.After(5 * time.Second):
		log.Println("request processed")
		w.Write([]byte("request processed"))
	case <-ctx.Done():
		log.Println("request cancelada pelo usuário")
		http.Error(w, "Error processing the request", http.StatusInternalServerError)
	}

}
