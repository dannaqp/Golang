package main

import (
	"fmt"
)

type Producto struct {
	Nombre        string
	Precio        float64
	CantidadDispo int
}

func main() {
	fmt.Println("Ejercicio 2: Gestión básica de productos usando Stuct")
	p1 := Producto{"Camiseta", 23.99, 15}
	p2 := Producto{"Pantalon", 45.00, 18}
	p3 := Producto{"Chaqueta", 50.99, 5}
	p4 := Producto{"Falda", 35.00, 0}
	var op int
	foriu := true
	for foriu {
		fmt.Println("Menú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Ver la información de un producto.\n Opción 2: Calcular el valor total del producto.\n Opción 3: Saber el stock de un producto. \n Opción 4: Ver el detalle de todos los productos.\nEscriba 0 para salir del menú ")
		fmt.Println("Escribe una opción: ")
		fmt.Scanln(&op)
		switch {
		case op == 0:
			fmt.Println("Salió del menú")
			foriu = false
		case op == 1:
			var pro int
			fmt.Println("Selecciona el producto del cual quieres ver el detalle: ", "\n1.", p1.Nombre, "\n2.", p2.Nombre, "\n3.", p3.Nombre, "\n4.", p4.Nombre)
			fmt.Println("Escribe una opción: ")
			fmt.Scanln(&pro)
			switch {
			case pro == 1:
				fmt.Println("Detalle del producto nro", pro)
				Info(p1)
			case pro == 2:
				fmt.Println("Detalle del producto nro", pro)
				Info(p2)
			case pro == 3:
				fmt.Println("Detalle del producto nro", pro)
				Info(p3)
			case pro == 4:
				fmt.Println("Detalle del producto nro", pro)
				Info(p4)
			}
			foriu = false
		case op == 2:
			var pro int
			fmt.Println("Selecciona el producto del cual quieres ver el valor total: ", "\n1.", p1.Nombre, "\n2.", p2.Nombre, "\n3.", p3.Nombre, "\n4.", p4.Nombre)
			fmt.Println("Escribe una opción: ")
			fmt.Scanln(&pro)
			switch {
			case pro == 1:
				fmt.Printf("\nEl valor total del producto %s es de $ %.2f", p1.Nombre, Total(p1))
			case pro == 2:
				fmt.Printf("\nEl valor total del producto %s es de $ %.2f", p2.Nombre, Total(p2))
			case pro == 3:
				fmt.Printf("\nEl valor total del producto %s es de $ %.2f", p3.Nombre, Total(p3))
			case pro == 4:
				fmt.Printf("\nEl valor total del producto %s es de $ %.2f", p4.Nombre, Total(p4))
			}
			foriu = false
		case op == 3:
			var pro int
			fmt.Println("Selecciona el producto del cual quieres conocer el stock: ", "\n1.", p1.Nombre, "\n2.", p2.Nombre, "\n3.", p3.Nombre, "\n4.", p4.Nombre)
			fmt.Println("Escribe una opción: ")
			fmt.Scanln(&pro)
			switch {
			case pro == 1:
				estado := Stock(p1)
				fmt.Println("El producto ", p1.Nombre, "tiene un stock de", p1.CantidadDispo, "\nEstado: ", estado)
			case pro == 2:
				estado := Stock(p2)
				fmt.Println("El producto ", p2.Nombre, "tiene un stock de", p2.CantidadDispo, "\nEstado: ", estado)
			case pro == 3:
				estado := Stock(p3)
				fmt.Println("El producto ", p3.Nombre, "tiene un stock de", p3.CantidadDispo, "\nEstado: ", estado)
			case pro == 4:
				estado := Stock(p4)
				fmt.Println("El producto ", p4.Nombre, "tiene un stock de", p4.CantidadDispo, "\nEstado: ", estado)
			}
			foriu = false
		case op == 4:
			fmt.Println("\tTienda de ropa")
			fmt.Printf("%-12s %-10s %-10s %-15s %s\n", "Nombre", "Precio", "Cantidad", "Valor Total", "Stock")
			fmt.Printf("%-12s $%-10.2f %-10d $%-15.2f %s\n", p1.Nombre, p1.Precio, p1.CantidadDispo, Total(p1), Stock(p1))
			fmt.Printf("%-12s $%-10.2f %-10d $%-15.2f %s\n", p2.Nombre, p2.Precio, p2.CantidadDispo, Total(p2), Stock(p2))
			fmt.Printf("%-12s $%-10.2f %-10d $%-15.2f %s\n", p3.Nombre, p3.Precio, p3.CantidadDispo, Total(p3), Stock(p3))
			fmt.Printf("%-12s $%-10.2f %-10d $%-15.2f %s\n", p4.Nombre, p4.Precio, p4.CantidadDispo, Total(p4), Stock(p4))
			foriu = false
		}
	}
}

func Info(p Producto) {
	fmt.Println("Nombre: ", p.Nombre)
	fmt.Println("Precio: ", p.Precio)
	fmt.Println("Cantidad: ", p.CantidadDispo)
}

func Total(p Producto) float64 {
	total := float64(p.CantidadDispo) * p.Precio
	return total
}

func Stock(p Producto) string {
	var dispo string
	if p.CantidadDispo > 0 {
		dispo = "Disponible"
	} else {
		dispo = "Sin Stock"
	}
	return dispo
}
