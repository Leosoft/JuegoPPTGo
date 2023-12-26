package main

import (
	"JuegoPPT/logica"
	_ "JuegoPPT/logica"
	"fmt"
	_ "fmt"
	"math/rand"
	_ "math/rand"
	_ "time"
)

func main() {

	opcion := ""

	for opcion != "3" {
		fmt.Print("Ingresa la opcion(1 para juego, 2 para salir): ")
		fmt.Scanln(&opcion)

		switch opcion {
		case "1":
			fmt.Println("Bienvenido")
			opciones := []string{"piedra", "papel", "tijera"}
			fmt.Print("Elige una opcion (piedra,papel,tijera): ")
			var usuario string
			fmt.Scanln(&usuario)

			if contiene(opciones, usuario) {
				computadora := opciones[rand.Intn(len(opciones))]
				logica.Juego(usuario, computadora)

			} else {
				fmt.Println("Opcion invalida, vuelve a intentarlo")
			}
		case "2":
			fmt.Println("Saliendo del programa, adios!")
			return

		default:
			fmt.Println("Vuelve a elegir")

		}
	}
}

// funcion auxiliar para verificar si un elemento esta en una lista de strings
func contiene(opciones []string, elemento string) bool {
	for _, opcion := range opciones {
		if opcion == elemento {
			return true
		}
	}
	return false
}
