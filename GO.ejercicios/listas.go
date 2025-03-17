package main

import (
	"fmt"
)

func main() {
	// Crear la lista
	lista := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	// Imprimir el elemento en la posición 4 (comienza en 0)
	fmt.Println(lista[4])

	// Imprimir la lista completa
	for _, dia := range lista {
		fmt.Print(dia, " ")
	}
	fmt.Println()

	// Sublista: Elementos desde la posición 0 hasta la 3
	for i := 0; i < 3; i++ {
		fmt.Print(lista[i], " ")
	}
	fmt.Println()

	// Lista con diferentes tipos de datos (usando interface{})
	var listaMixta []interface{}
	listaMixta = append(listaMixta, "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", 1, 2, 3)

	// Crear una lista interna dentro de la lista principal
	datosEsteban := []interface{}{"Esteban", 0.2, 2.25, true}

	// Agregar la lista interna a la lista principal
	listaMixta = append(listaMixta, datosEsteban)

	// Imprimir los primeros 4 elementos
	for i := 0; i < 4; i++ {
		switch v := listaMixta[i].(type) {
		case string:
			fmt.Print(v, " ")
		case int:
			fmt.Print(v, " ")
		case float64:
			fmt.Print(v, " ")
		case bool:
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	// Imprimir los primeros 3 elementos de la lista interna
	if subLista, ok := listaMixta[9].([]interface{}); ok {
		for i := 0; i < 3; i++ {
			switch v := subLista[i].(type) {
			case string:
				fmt.Print(v, " ")
			case int:
				fmt.Print(v, " ")
			case float64:
				fmt.Print(v, " ")
			case bool:
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}
}