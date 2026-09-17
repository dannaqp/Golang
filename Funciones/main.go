package main

import "fmt"

func saludar() {
	fmt.Println("Hola esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func suma(a, b int) int {
	return a + b
}

func main() {

	var usr string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	var a, b int
	fmt.Println("Ingresa dos valores: ")
	fmt.Scan(&a, &b)
	fmt.Println("El resultado es: ", suma(a, b))
}
