package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8081", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
