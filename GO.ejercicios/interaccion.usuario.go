package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variables para almacenar los datos del usuario
	var nombres, apellidos, profesion string
	var anioNacimiento int

	// Pedir al usuario que ingrese sus datos
	fmt.Print("Escriba sus nombres completos: ")
	fmt.Scanln(&nombres)

	fmt.Print("Escriba sus Apellidos completos: ")
	fmt.Scanln(&apellidos)

	fmt.Print("Escriba su profesion: ")
	fmt.Scanln(&profesion)

	fmt.Print("Escriba su año de nacimiento: ")
	fmt.Scanln(&anioNacimiento)

	// Calcular la edad en base al año actual (2025)
	edad := 2025 - anioNacimiento

	// Mostrar el resultado
	fmt.Printf("El (La) %s %s %s tiene %d años\n", profesion, nombres, apellidos, edad)
}