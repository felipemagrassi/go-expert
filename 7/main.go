package main

import "fmt"

type Address struct {
	City, State, Street string
	Number              int
}

// Similar to implements in Java
type Person interface {
	changeName(name string)
}

type Seller struct {
	Name string
}

type Client struct {
	Name        string
	Active      bool
	Age         int
	MainAddress Address
}

func newName(p Person, name string) {
	p.changeName(name)
}

func (c Client) changeName(name string) {
	c.Name = name
}

func (c Client) Deactivate() Client {
	c.Active = false

	return c
}

func main() {
	address := Address{"São Paulo", "SP", "Rua dos Bobos", 0}
	felipe := Client{Name: "Felipe", Active: true, Age: 23, MainAddress: address}

	fmt.Printf("%+v\n", felipe)
	fmt.Printf("Name: %s, Active: %t, Age: %d\n", felipe.Name, felipe.Active, felipe.Age)
	fmt.Printf("Address: %s, %s, %s, %d\n", felipe.MainAddress.City, felipe.MainAddress.State, felipe.MainAddress.Street, felipe.MainAddress.Number)

	felipe = felipe.Deactivate()

	fmt.Printf("Name: %s, Active: %t, Age: %d\n", felipe.Name, felipe.Active, felipe.Age)

}
