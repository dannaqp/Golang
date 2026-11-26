package main

import "fmt"

func calcularCuadrado(num int, canal chan int) {
	res := num * num
	canal <- res
}

func main() {
	canal := make(chan int)
	go calcularCuadrado(4, canal)

	resultado := <-canal

	fmt.Println("El resultado es: ", resultado)
}
