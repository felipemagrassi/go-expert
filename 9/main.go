package main

import "fmt"

type Client struct {
	name    string
	balance float64
}

func (c Client) changeName(name string) {
	c.name = name
}

func (c *Client) changeNamePointer(name string) {
	c.name = name
}

func newClient(name string) *Client {
	return &Client{name: name, balance: 0}
}

func main() {
	felipe := newClient("Felipe")
	fmt.Printf("%T\n", felipe)
	fmt.Println(felipe)
	fmt.Println("Name: ", felipe.name)
	felipe.changeName("Felipe 2")
	fmt.Println("Name: ", felipe.name)
	felipe.changeNamePointer("Felipe 3")
	fmt.Println("Name: ", felipe.name)
}
