package main

import "fmt"

func main() {
	canal := make(chan string)
	canal <- "Olá Mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}