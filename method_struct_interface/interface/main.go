package main

import "fmt"

type Carro interface {
	acelerar()
}

type Van struct{}

func (v *Van) acelerar() {
	fmt.Println("Van acelerando")
}

func main() {
	van := Van{}
	printarAceleracao(&van)
}

func printarAceleracao(c Carro) {
	c.acelerar()
}
