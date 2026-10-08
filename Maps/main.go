package main

import "fmt"

func main() {
	votos := RegistrarVotos()

	fmt.Println("\nResultados de la votación:")
	for actividad, cantidad := range votos {
		fmt.Printf("  - %s: %d votos\n", actividad, cantidad)
	}

	actividadGanadora, votosGanadores := DeterminarGanador(votos)
	fmt.Printf("\nLa actividad más votada es: %s (%d votos)\n", actividadGanadora, votosGanadores)
}

func RegistrarVotos() map[string]int {
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"música":      0,
	}

	fmt.Println("Sistema de Votación")
	fmt.Println("Opciones: deportes, videojuegos, cine, música")

	var voto string
	for i := 0; i < 5; i++ {
		fmt.Printf("Ingrese el voto %d: ", i+1)
		fmt.Scanln(&voto)

		_, existe := votos[voto]
		if existe {
			votos[voto]++
		}
	}
	return votos
}

func DeterminarGanador(votos map[string]int) (string, int) {
	maxVotos := -1
	ganador := ""

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			ganador = actividad
		}
	}
	return ganador, maxVotos
}
