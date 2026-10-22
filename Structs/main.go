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

	if product2.Vender(5) == true {
		fmt.Println("El stock restante despues de la venta de 5", product2.Nombre, "es: ", product2.stock)
	}
}

func calculartotal(p Producto) float64 {
	total := p.Precio * float64(p.stock)
	return total
}

func (p *Producto) Vender(cantidad int64) bool {
	if cantidad < 0 || cantidad > p.stock {
		return false
	}
	p.stock -= cantidad
	return true
}
