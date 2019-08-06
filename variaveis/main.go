package main

import "fmt"

func main() {
	// será alocado o espaço necessário para a
	// declaração da variável
	var variavelExplicita1 string
	variavelExplicita1 = "olá"

	// especificar a tipagem pode ser interessante pq,
	// por exemplo, ao declarar um int pequeno, será
	// alocado um int64 desnecessariamente
	var variavelExplicita2 int8 = 2

	// forma mais bonita de declarar uma variável
	// (possível somente dentro de função)
	shorDescription := "Valor"

	var primeiro, segundo int = 1, 2

	var (
		texto    = "oi"
		numero   = 2
		booleano = true
	)

	fmt.Println(variavelExplicita1, shorDescription, primeiro, segundo)
	fmt.Println(variavelExplicita2, texto, numero, booleano)
}
