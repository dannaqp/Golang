package main

import "fmt"

type Persona struct {
	nombre string
}

func (p Persona) GetNombre() string {
	return p.nombre
}
func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	p1 := Persona{}
	p1.SetNombre("Danna")

	fmt.Println("Su nombre es: ", p1.GetNombre())
	p1.SetNombre("Camu")
	fmt.Println("Ahora su nombre es: ", p1.GetNombre())
}
