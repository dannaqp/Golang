package main

import (
	"fmt"
	"taller2/contador"
	"taller2/conversor"
)

func main() {
	fmt.Println("-----Taller Práctico Paquetes-----")
	var op int
	foriu := true
	for foriu {
		fmt.Println("Menú Principal")
		fmt.Println("Selecciona una opción: \n Opción 1: Conversor de monedas \n Opción 2: Contador de Vocales\n Escriba 0 para salir del menú ")

		fmt.Println("Escribe una opción: ")
		fmt.Scan(&op)
		switch {
		case op == 0:
			fmt.Println("Salió del menú")
			foriu = false
		case op == 1:
			fmt.Println("Conversor de monedas")
			var dlr float64
			var mon int
			fmt.Println("Ingrese el valor en dólares que desear convertir: ")
			fmt.Scan(&dlr)
			fmt.Println("¿A qué moneda desea convertir?")
			fmt.Println("Selecciona una opción: \n Opción 1: Euros \n Opción 2: LB (Libras Esterlinas)\n Opción 3: Won (Sur Koreano)\n Opción 4: BTC (Bitcon)")
			fmt.Scan(&mon)
			switch {
			case mon == 1:
				result := conversor.Eur(dlr)
				fmt.Println("El equivalente en Euros es: ", result)
				foriu = false
			case mon == 2:
				result := conversor.Lb(dlr)
				fmt.Println("El equivalente en Libras Esterlinas es: ", result)
				foriu = false
			case mon == 3:
				result := conversor.Won(dlr)
				fmt.Println("El equivalente en Wones es: ", result)
				foriu = false
			case mon == 4:
				result := conversor.Btc(dlr)
				fmt.Println("El equivalente en Bitcon es: ", result)
				foriu = false
			}
		case op == 2:
			fmt.Println("Contador de Vocales")
			var or string
			fmt.Println("Ingrese la frase de la cual desea contar las vocales: ")
			fmt.Scan(&or)
			fmt.Println("El total de veces que aparece cada vocal es: ")
			fmt.Println("Vocal a: ", contador.A(or))
			fmt.Println("Vocal e: ", contador.E(or))
			fmt.Println("Vocal i: ", contador.I(or))
			fmt.Println("Vocal o: ", contador.O(or))
			fmt.Println("Vocal u: ", contador.U(or))
		}
	}
}
