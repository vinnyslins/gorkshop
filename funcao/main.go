package main

import "fmt"

func main() {
	falar("oi", printar)
	texto, numero := retornoMultiplo("oi", 1)
	fmt.Println(texto, numero)
}

func soma(a, b int) int {
	return a + b
}

func somaNomeada(a, b int) (resultado int) {
	resultado = a + b
	return
}

func falar(text string, funcao func(string)) {
	funcao(text)
}

func printar(text string) {
	fmt.Println(text)
}

func retornoMultiplo(a string, b int) (string, int) {
	return a, b
}
