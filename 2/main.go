package main

const t = true

type ID int

var (
	a bool    = true
	b bool    = false
	c string  = "hello"
	d int     = 1
	e float64 = 1.2
	g ID      = 1
)

func main() {
	// short hand declaration
	f := true

	// f := false // error: no new variables on left side of :=
	println(a, b, c, d, e, f, g, t)

}
