package main

import "fmt"

func main() {
	notas := AlmacenarNotas()
	promedios := make([]float64, 6)
	max := make([]float64, 6)
	min := make([]float64, 6)
	for i := 0; i < 6; i++ {
		notasIndiv := notas[i][:]

		promedios[i] = PromIndiv(notasIndiv)
		max[i] = Mayor(notasIndiv)
		min[i] = Menor(notasIndiv)
	}
	promedioGeneral := PromGen(promedios)

	fmt.Println("Análisis de Notas")
	for i := 0; i < 6; i++ {
		fmt.Printf("Estudiante Nro %d:\n", i+1)
		fmt.Print(" ")
		fmt.Printf("  - Promedio: %.2f\n", promedios[i])
		fmt.Print(" ")
		fmt.Printf("  - Nota más Alta: %.2f\n", max[i])
		fmt.Print(" ")
		fmt.Printf("  - Nota más Baja: %.2f\n", min[i])
		fmt.Print(" ")
	}

	fmt.Printf("Promedio general de la clase: %.2f \n", promedioGeneral)
}

func AlmacenarNotas() [6][4]float64 {
	var notas = [6][4]float64{}
	for i := 0; i < 6; i++ {
		fmt.Println("Ingresa las notas del estudiante nro: ", i+1)
		fmt.Println("Nota de matemáticas: ")
		fmt.Scanln(&notas[i][0])
		fmt.Println("Nota de artes: ")
		fmt.Scanln(&notas[i][1])
		fmt.Println("Nota de lenguaje: ")
		fmt.Scanln(&notas[i][2])
		fmt.Println("Nota de física: ")
		fmt.Scanln(&notas[i][3])
	}
	return notas
}

func PromIndiv(NotasEst []float64) float64 {
	suma := 0.0
	for _, nota := range NotasEst {
		suma += nota
	}
	return suma / float64(len(NotasEst))
}

func Mayor(NotasEst []float64) float64 {
	max := NotasEst[0]
	for i := 1; i < len(NotasEst); i++ {
		if NotasEst[i] > max {
			max = NotasEst[i]
		}
	}
	return max
}

func Menor(NotasEst []float64) float64 {
	min := NotasEst[0]
	for i := 1; i < len(NotasEst); i++ {
		if NotasEst[i] < min {
			min = NotasEst[i]
		}
	}
	return min
}

func PromGen(promedios []float64) float64 {
	suma := 0.0
	for _, prom := range promedios {
		suma += prom
	}
	return suma / float64(len(promedios))
}
