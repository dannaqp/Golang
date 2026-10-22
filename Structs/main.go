package main

import "fmt"

type Producto struct {
	Nombre string
	Precio float64
	stock  int64
}

func main() {

	product := Producto{}

	product.Nombre = "Intel Core Ultra 7"
	product.Precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		Precio: 30.5,
		stock:  15,
	}
	fmt.Println(product)
	fmt.Println(product2)

	fmt.Println("El Producto", product2.Nombre, "tiene un valor total de inventario de: ", calculartotal(product2))
}

func calculartotal(p Producto) float64 {
	total := p.Precio * float64(p.stock)
	return total
}
