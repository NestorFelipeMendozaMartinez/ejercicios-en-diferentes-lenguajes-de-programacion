package main

import (
	"fmt"
)

func main() {
	// Solicitar y leer los valores de A y B
	var A, B int
	fmt.Print("Ingrese valor para A: ")
	fmt.Scan(&A)
	fmt.Print("Ingrese valor para B: ")
	fmt.Scan(&B)

	// Realizar las operaciones
	suma := A + B
	fmt.Println("La suma de los números es:", suma)

	res := A - B
	fmt.Println("La resta de los números es:", res)

	multi := A * B
	fmt.Println("La multiplicación de los números es:", multi)

	// Se asume que la división debe manejarse para evitar división por 0
	if B != 0 {
		div := float64(A) / float64(B)
		fmt.Println("La división de los números es:", div)
	} else {
		fmt.Println("Error: División por cero no permitida")
	}

	// Comparaciones
	igual := A == B
	fmt.Println("¿El número es igual?", igual)

	mayor := A > B
	fmt.Println("¿El número mayor es A?", mayor)

	// Mostrar el número menor y mayor
	if A > B {
		fmt.Println("El número mayor es:", A)
		fmt.Println("El número menor es:", B)
	} else {
		fmt.Println("El número mayor es:", B)
		fmt.Println("El número menor es:", A)
	}
}