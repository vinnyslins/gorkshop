package main

import (
	"fmt"
)

func main() {
	go c()
	go a()

	// time.Sleep(time.Millisecond * 2)
	// fmt.Println("Eu sou o principal")
}

func c() {
	fmt.Println("Estou na concorrêcia")
}

func a() {
	fmt.Println("Eu também")
}
