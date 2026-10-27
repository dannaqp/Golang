package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func main() {
	inventario := Inventario{
		Productos: make(map[int]Producto),
	}

	fmt.Println("Sistema de Gestión de Inventario")
	var op int
	foriu := true

	for foriu {
		fmt.Println("\nMenú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Agregar producto \n Opción 2: Listar productos\n Opción 3: Buscar producto\n Opción 4: Actualizar producto\n Opción 5: Eliminar producto\n Escriba 0 para salir del menú ")

		fmt.Println("Escribe una opción: ")
		fmt.Scan(&op)

		switch {
		case op == 0:
			fmt.Println("Salió del menú")
			foriu = false

		case op == 1:
			fmt.Println("Agregar producto")
			var id, cantidad int
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

			producto := NuevoProducto(id, nombre, precio, cantidad)
			inventario.AgregarProducto(producto)
			fmt.Println("Producto agregado con éxito")
			foriu = false

		case op == 2:
			fmt.Println("Lista de productos")
			if len(inventario.Productos) == 0 {
				fmt.Println("El inventario está vacío.")
			} else {
				inventario.ListarProductos()
			}
			foriu = false

		case op == 3:
			fmt.Println("Buscar producto")
			var id int
			fmt.Print("Ingrese el ID del producto que desea buscar: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)
			if existe {
				producto.MostrarInfo()
			} else {
				fmt.Println("Producto no encontrado.")
			}
			foriu = false

		case op == 4:
			fmt.Println("Actualizar producto")
			var id, cantidad int
			var precio float64
			fmt.Print("Ingrese el ID del producto que desea actualizar: ")
			fmt.Scan(&id)

			producto, existe := inventario.BuscarProducto(id)
			if existe {
				fmt.Print("Ingrese el nuevo precio: ")
				fmt.Scan(&precio)
				fmt.Print("Ingrese la nueva cantidad: ")
				fmt.Scan(&cantidad)

				producto.ActualizarPrecio(precio)
				producto.ActualizarCantidad(cantidad)
				inventario.AgregarProducto(producto)

				fmt.Println("¡Producto actualizado!")
			} else {
				fmt.Println("Producto no encontrado.")
			}
			foriu = false

		case op == 5:
			fmt.Println("Eliminar producto")
			var id int
			fmt.Print("Ingrese el ID del producto que desea eliminar: ")
			fmt.Scan(&id)

			inventario.EliminarProducto(id)
			fmt.Println("Inventario actualizado:")
			inventario.ListarProductos()
			foriu = false
		}
	}
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{id, nombre, precio, cantidad}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)
	fmt.Println("Cantidad:", p.Cantidad)
}

type Inventario struct {
	Productos map[int]Producto
}

func (i *Inventario) AgregarProducto(p Producto) {
	i.Productos[p.ID] = p
}

func (i *Inventario) EliminarProducto(id int) {
	delete(i.Productos, id)
}

func (i *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := i.Productos[id]
	return p, existe
}

func (i *Inventario) ListarProductos() {
	for _, p := range i.Productos {
		p.MostrarInfo()
	}
}
