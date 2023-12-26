package logica

import (
	"fmt"
	_ "fmt"
)

func Juego(usuario, computadora string) {
	fmt.Printf("Tu elegiste: %s\n", usuario)
	fmt.Printf("La computadora eligio: %s\n", computadora)

	if usuario == computadora {
		fmt.Println("Es un empate")
	} else if (usuario == "piedra" && computadora == "tijera") || (usuario == "papel" && computadora == "piedra") ||
		(usuario == "tijera" && computadora == "papel") {
		fmt.Println("GANASTE!!!")
	} else {
		fmt.Println("La computadora gana")
	}
}
