package main

import "fmt"

func main() {

	c := new(children)
	fmt.Println(c.GetMessage())

}

type base interface {
	GetMessage() string
	GetName() string
}

type parent struct{}

func (p *parent) GetName() string {
	return "Ivan"
}

func (p *parent) GetMessage() string {
	return p.getMessage(p)
}

func (p *parent) getMessage(b base) string {
	return b.GetName()
}

type children struct {
	parent
}

func (c *children) GetMessage() string {
	return c.getMessage(c)
}

func (c *children) GetName() string {
	return "John"
}
