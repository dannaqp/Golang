package main

import (
	"errors"
	"fmt"
)

// Creacion de la interface

type Dispositivo interface {
	Encender() error
	Apagar() error
	Estado() string
}

// Creacion de las estructuras

type Servidor struct {
	Nombre string
	Activo bool
}

type Router struct {
	Nombre    string
	Encendido bool
}

// Funciones de Servidor

func (s *Servidor) Encender() error {
	if s.Activo {
		return errors.New("El servidor se encuentra activo")
	}
	s.Activo = true
	return nil
}

func (s *Servidor) Apagar() error {
	if !s.Activo {
		return errors.New("El servidor no se encuentra activo")
	}
	s.Activo = false
	return nil
}

func (s *Servidor) Estado() string {
	if s.Activo {
		return s.Nombre + ":" + "se encuentra activo"
	}
	return s.Nombre + ":" + "no se encuentra activo"
}

// Funciones de Router

func (r *Router) Encender() error {
	if r.Encendido {
		return errors.New("El router ya esta encendido")
	}
	r.Encendido = true
	return nil
}

func (r *Router) Apagar() error {
	if !r.Encendido {
		return errors.New("El router ya esta apagado")
	}
	r.Encendido = false
	return nil
}

func (r *Router) Estado() string {
	if r.Encendido {
		return r.Nombre + ":" + "esta encendido"
	}
	return r.Nombre + ":" + "esta apagado"
}

func OperarDispositivo(d Dispositivo) {
	fmt.Println("Estatus Inicial: ", d.Estado())
	err := d.Encender()

	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("Despues de haber encendido: ", d.Estado())

	er := d.Apagar()

	if er != nil {
		fmt.Println("ERROR: ", er)
	}
	fmt.Println("Despues de haber apagado: ", d.Estado())
}

func main() {
	router := &Router{
		Nombre:    "Cisco",
		Encendido: true,
	}
	servidor := &Servidor{
		Nombre: "Servidor Web",
	}
	fmt.Println("ROUTER")
	OperarDispositivo(router)

	fmt.Println("SERVIDOR")
	OperarDispositivo(servidor)
}
