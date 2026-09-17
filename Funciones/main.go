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

func sumaresta(c, d int) (int, int) {
	if c > d {
		return c + d, c - d
	}
	return c + d, 0
}

func mostrarNum(num ...int) {
	fmt.Println("Los números ingresados son: ", num)
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
	// Solo para un retorno
	fmt.Println("El resultado es: ", suma(a, b))
	var c, d int
	fmt.Println("Ingresa dos valores: ")
	fmt.Scan(&c, &d)

	w, z := sumaresta(c, d)
	fmt.Println("El resultado de la suma es: ", w, "y de la resta es: ", z)

	mostrarNum(5, 10, 15, 20, 25)

	fmt.Println("La sumatoria es: ", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sumatoria(nums ...int) int {
	total := 0
	for _, nro := range nums {
		total += nro
	}
	return total
}
