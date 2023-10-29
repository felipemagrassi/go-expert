package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Golang", 500}

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso {{.Nome}} - CH: {{.CargaHoraria}}")

	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
