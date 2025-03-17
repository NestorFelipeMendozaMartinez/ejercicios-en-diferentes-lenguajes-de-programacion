package main

import (
	"fmt"
	"math" // Importa el paquete math para usar la función Pow
)

func main() {
	// Solicitar el primer valor
	fmt.Println("Ingrese el primer valor (base):")
	var A int
	fmt.Scanln(&A)

	// Inicializar el valor de B (en el código original no se usa, pero se podría mantener si lo deseas)
	var B int = 0

	// Solicitar el segundo valor (exponente)
	fmt.Println("Ingrese el segundo valor (exponente):")
	var C int
	fmt.Scanln(&C)

	// Calcular la potencia
	// math.Pow devuelve un tipo float64, por lo que necesitamos hacer la conversión a int
	valor := int(math.Pow(float64(A), float64(C)))

	// Mostrar el resultado
	fmt.Printf("La potencia de %d sobre %d es: %d\n", A, C, valor)
}