package main

import "fmt"

func main() {
	variavel := 1

	switch variavel {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("nenhum")
	}
}
