package main

import (
	"errors"
	"fmt"
)

type Estudiante struct {
	edad int
}

func (e Estudiante) GetEdad() int {
	return e.edad
}

func (e *Estudiante) SetEdad(edad int) error {
	if edad < 0 || edad > 80 {
		return errors.New("ERROR, La edad  debe estar entre 1 y 79 aos!!")
	}
	e.edad = edad
	return nil
}

func NewEstudiante(edadin int) (*Estudiante, error) {
	est := &Estudiante{}

	err := est.SetEdad(edadin)

	if err != nil {
		return nil, err
	}
	return est, nil
}

func main() {
	var edadin int
	fmt.Print("Ingrese la nueva edad: ")
	fmt.Scanln(&edadin)

	nuevoEst, errorr := NewEstudiante(edadin)

	if errorr != nil {
		fmt.Println("No se pudo crear el estudiante:", errorr)
	} else {
		fmt.Println("Estudiante creado exitosamente. Ahora la edad es: ", nuevoEst.GetEdad())
	}
}
