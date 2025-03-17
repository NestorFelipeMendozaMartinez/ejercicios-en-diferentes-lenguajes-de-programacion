package main

import (
	"fmt"
)

func main() {
	// Variables para almacenar voltaje y resistencia
	var voltaje, resistencia int

	// Pedir al usuario que ingrese los valores de voltaje y resistencia
	fmt.Print("Ingrese el valor del voltaje del circuito: ")
	fmt.Scan(&voltaje)

	fmt.Print("Ingrese el valor de la resistencia del circuito: ")
	fmt.Scan(&resistencia)

	// Calcular la intensidad (amperaje)
	intensidad := float64(voltaje) / float64(resistencia)

	// Mostrar el resultado
	fmt.Printf("Al conectar un resistor de R %d ohm a una fuente de V %d voltaje circulara una corriente de %.2f amperios\n", resistencia, voltaje, intensidad)
}