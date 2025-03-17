package main

import "fmt"

func main() {
	// Definir las variables booleanas
	a := true
	b := false

	// Operación lógica 'and'
	fmt.Println(a && b)

	// Definir variables numéricas
	x := 2
	y := 3
	z := 4
	w := 10

	// Operación lógica con comparación '==', '<' y '>'
	fmt.Println(x == y && z < w)

	// Operación lógica con 'not' (invertir una condición)
	fmt.Println(!(x == y) && z > w)
}