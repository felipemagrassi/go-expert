package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("content.html").ParseFiles(templates...))

		err := tmp.Execute(w, Cursos{{"Golang", 500}, {"Python", 123}, {"NodeJS", 500}})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)

}
