package main

import "fmt"

func main() {
	c := make(chan string)

	go func() { c <- "Hello" }()
	go func() { c <- "World" }()

	msg1 := <-c
	msg2 := <-c

	fmt.Println(msg1, msg2)
}
