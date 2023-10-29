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
	cursos := []Curso{{"Golang", 500}, {"Python", 123}, {"NodeJS", 500}}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmp.Execute(w, cursos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)

}
