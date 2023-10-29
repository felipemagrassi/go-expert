package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	curso := Curso{"Golang", 500}
	curso2 := Curso{"Python", 400}

	file, err := os.Create("index.html")

	defer file.Close()

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err = tmp.Execute(file, Cursos{curso, curso2})
	if err != nil {
		panic(err)
	}

}
