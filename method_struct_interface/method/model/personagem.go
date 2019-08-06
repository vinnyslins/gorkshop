package model

import "fmt"

// Personagem alguma coisa (isso é uma doc)
type Personagem struct {
	name   string
	age    int
	gender string
}

// SetName é uma função que atribuí o name
func (p *Personagem) SetName(name string) {
	p.name = name
}

// Falar é uma função que printa o name
func (p *Personagem) Falar() {
	fmt.Println("Meu nome é", p.name)
}
