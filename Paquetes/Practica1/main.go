package main

import (
	"Practica/operaciones"
	"Practica/saludo"
	"fmt"
)

func main() {
	fmt.Println("****Bienvenid@s a la Clase de Paquetes****")
	mensaje := saludo.Saludar("Camu")
	fmt.Println(mensaje)
	sum, rest := operaciones.Sumaresta(18, 11)
	fmt.Println("El resultado de suma es: ", sum, " y de la resta es: ", rest)
}
