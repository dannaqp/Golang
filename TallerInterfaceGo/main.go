package main

import (
	"fmt"
	"time"
)

type InformacionProducto interface {
	MostrarInfo() string
}

type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p ProductoBase) MostrarInfo() string {
	return fmt.Sprintf("ID: %d\nNombre: %s\nPrecio: %.2f\nCantidad: %d", p.ID, p.Nombre, p.Precio, p.Cantidad)
}

func (p *ProductoBase) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *ProductoBase) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

type Electronico struct {
	ProductoBase
	GarantiaMeses int
}

func (e Electronico) MostrarInfo() string {
	return fmt.Sprintf("%s\nGarantía: %d meses", e.ProductoBase.MostrarInfo(), e.GarantiaMeses)
}

func NuevoProductoElectronico(id int, nombre string, precio float64, cantidad int, garantia int) Electronico {
	return Electronico{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		GarantiaMeses: garantia,
	}
}

type Alimento struct {
	ProductoBase
	FechaExpiracion time.Time
}

func (a Alimento) MostrarInfo() string {
	return fmt.Sprintf("%s\nFecha de Exp: %s", a.ProductoBase.MostrarInfo(), a.FechaExpiracion.Format("2006-01-02"))
}

func NuevoProductoAlimento(id int, nombre string, precio float64, cantidad int, fechaExp time.Time) Alimento {
	return Alimento{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		FechaExpiracion: fechaExp,
	}
}

type Inventario struct {
	Productos []InformacionProducto
}

func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}

func (i *Inventario) ListarProductos() {
	if len(i.Productos) == 0 {
		fmt.Println("El inventario se encuentra vacio")
		return
	}
	for _, p := range i.Productos {
		fmt.Println(p.MostrarInfo())
	}
}

func main() {
	inventario := Inventario{}

	fmt.Println("Sistema de Gestión de Inventario")
	var op int
	foriu := true

	for foriu {
		fmt.Println("\nMenú Principal")
		fmt.Println("Selecciona una opción: \nOpción 1: Agregar producto Electrónico\nOpción 2: Agregar producto Alimento\nOpción 3: Listar productos\nEscriba 0 para salir del menú")

		fmt.Print("Escribe una opción: ")
		fmt.Scan(&op)

		switch {
		case op == 0:
			fmt.Println("Salió del menú")
			foriu = false

		case op == 1:
			fmt.Println("Agregar Electronico")
			var id, cantidad, garantia int
			var nombre string
			var precio float64

			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)
			fmt.Print("Ingrese el nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingrese el precio: ")
			fmt.Scan(&precio)
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scan(&cantidad)
			fmt.Print("Ingrese los meses de garantia del electronico: ")
			fmt.Scan(&garantia)

			producto := NuevoProductoElectronico(id, nombre, precio, cantidad, garantia)
			inventario.AgregarProducto(producto)
			fmt.Println("El producto electronico ha sido agregado con exito")
			foriu = false

		case op == 2:
			fmt.Println("Agregar Alimento")
			var id, cantidad int
			var nombre, fechaStr string
			var precio float64

			fmt.Print("Ingrese el ID: ")
			fmt.Scan(&id)
			fmt.Print("Ingrese el nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Ingrese el precio: ")
			fmt.Scan(&precio)
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scan(&cantidad)
			fmt.Print("Ingrese la fecha de expiracion del alimento (AAAA-MM-DD): ")
			fmt.Scan(&fechaStr)

			fechaExp, err := time.Parse("2006-01-02", fechaStr)
			if err != nil {
				fmt.Println("Formato de fecha inválido, use AAAA-MM-DD (ej. 2026-12-31)")
				continue
			}

			producto := NuevoProductoAlimento(id, nombre, precio, cantidad, fechaExp)
			inventario.AgregarProducto(producto)
			fmt.Println("El alimento ha sido agregado con exito")
			foriu = false

		case op == 3:
			fmt.Println("Lista de productos")
			inventario.ListarProductos()
			foriu = false
		}
	}
}
