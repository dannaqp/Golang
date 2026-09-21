package main

import "fmt"

// Taller 1 Funciones

func main() {
	var op int
	foriu := true
	for foriu {
		fmt.Println("Menú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Notas \n Opción 2: Suma de números\n Opción 3: Conversión de Celsius a Fahrenheit \n Opción 4: Conversión de Fahrenheit a Celsius \n Escriba 0 para salir del menú ")

		fmt.Println("Escribe una opción: ")
		fmt.Scan(&op)
		switch {
		case op == 0:
			fmt.Println("Saliendo del programa...")
			foriu = false
		case op == 1:
			var est int
			var sum float64
			fmt.Println("Escriba el número de estudiantes de los que desea ingresar notas: ")
			fmt.Scan(&est)
			sum = averageGrade(notas(est), est)
			if sum >= 70 {
				fmt.Println("El promedio del curso es:", sum, ", es aprobado")
			} else {
				fmt.Println("El promedio del curso es:", sum, ", es reprobado")
			}
			switch {
			case sum >= 90 && sum <= 100:
				fmt.Println("Excellent performance")
			case sum >= 80 && sum <= 89:
				fmt.Println("Good performance")
			case sum >= 70 && sum <= 79:
				fmt.Println("Satisfactory performance")
			case sum < 70:
				fmt.Println("Needs improvement")
			}
		case op == 2:
			var nro int
			fmt.Println("Escriba el número para la suma: ")
			fmt.Scan(&nro)
			fmt.Println("La suma total desde 1 a ", nro, "es", sumanro(nro))

		case op == 3:
			var cel int
			fmt.Println("Ingrese la temperatura en Celsius: ")
			fmt.Scan(&cel)
			//celafah(cel)
		case op == 4:
			var fah int
			fmt.Println("Ingrese la temperatura en Fahrenheit: ")
			fmt.Scan(&fah)
			//fahacel(fah)
		}
	}

}

func notas(est int) float64 {
	var nota float64
	var sum float64
	for i := 0; i < est; i++ {
		fmt.Println("Ingrese la nota (0-100) del estudiante nro: ", (i + 1))
		fmt.Scan(&nota)
		if nota < 0 || nota > 100 {
			fmt.Println("La nota ingresada debe ser entre 0 y 100, por favor ingrese la nota nuevamente.")
			i--
			continue
		} else {
			sum += nota
		}

	}
	return sum
}

func averageGrade(sum float64, est int) float64 {
	return sum / float64(est)
}

func sumanro(nro int) int {
	total := 0
	for i := 1; i <= nro; i++ {
		total += i
	}
	return total

}

/*
func celafah(cel float64) {

}
func fahacel(fah float64) {

}
*/
