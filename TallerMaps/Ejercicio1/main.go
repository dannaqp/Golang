package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ejercicio 1: Conteo de palabras usando maps")
	contador := Separador()
	fmt.Println("El conteo de las palabras de tu frase es: ")
	for palabra, cont := range contador {
		fmt.Printf("\n%s: %d ", palabra, cont)
	}
	pal, rep := MasUs(contador)
	fmt.Printf("\nLa palabra más usada es '%s' con %d repeticiones.\n", pal, rep)

}

func Separador() map[string]int {
	fmt.Println("Ingresa tu frase:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	frase := sc.Text()
	palabras := strings.Fields(frase)
	contador := make(map[string]int)
	for _, pal := range palabras {
		pal2 := strings.ToLower(pal)
		palabra := strings.Trim(pal2, ".,¡!¿?();:\"'")
		contador[palabra]++
	}
	return contador
}

func MasUs(contador map[string]int) (string, int) {
	var palabrarep string
	var rep int
	for palabra, cantidad := range contador {
		if cantidad > rep {
			rep = cantidad
			palabrarep = palabra
		}
	}
	return palabrarep, rep
}
