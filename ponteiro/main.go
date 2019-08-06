package main

import "fmt"

func main() {
	valor := 10
	novoValor := &valor

	incrementar(&valor)
	fmt.Println(*novoValor)
}

func incrementar(numero *int) {
	*numero++
}
