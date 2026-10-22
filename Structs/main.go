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

	if product2.Vender(15) == true {
		fmt.Println("El producto", product2.Nombre, "aun tiene stock disponible, el restante es de: ", product2.stock)
	} else {
		fmt.Println("El producto", product2.Nombre, "Ya no tiene stock")
	}

	product3 := NewProduct("mouse", 12.91, 0)
	fmt.Println(product3)
	fmt.Println("El nuevo stock del producto", product3, "es: ", product3.AsignarStock(18))
}

func calculartotal(p Producto) float64 {
	total := p.Precio * float64(p.stock)
	return total
}

func (p *Producto) Vender(cantidad int64) bool {
	if cantidad < 0 || cantidad >= p.stock {
		return false
	} else {
		p.stock -= cantidad
		return true
	}
}

func NewProduct(nombre string, precio float64, stock int64) Producto {
	return Producto{Nombre: nombre, Precio: precio, stock: 0}
}

func (p *Producto) AsignarStock(newstock int64) int64 {
	if newstock > 0 {
		p.stock += newstock
	}
	return p.stock
}
