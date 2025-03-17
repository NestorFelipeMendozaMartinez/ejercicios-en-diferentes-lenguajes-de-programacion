package main

import "fmt"

// Función que imprime un texto
func mostrarTexto() {
	fmt.Println("Hola")
}

// Función que multiplica 5 por 8 e imprime el resultado
func multiplicar() {
	a := 5
	b := 8
	fmt.Println(a * b)
}

// Función que usa variables globales y las imprime
func multiplicarConVariablesGlobales(a int, b int) {
	fmt.Println(a * b)
}

// Función que devuelve el resultado de la multiplicación
func multiplicarConReturn() int {
	a := 5
	b := 8
	return a * b
}

// Función que valida si el valor de a es igual a 5
func validarDato(a int) bool {
	return a == 5
}

func main() {
	// Ejemplo de multiplicación de elementos de dos arrays
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	var c []int

	// Multiplicar los elementos correspondientes de a y b y almacenarlos en c
	for i := 0; i < 5; i++ {
		c = append(c, a[i]*b[i])
	}

	// Imprimir el slice c
	fmt.Print("[ ")
	for _, num := range c {
		fmt.Print(num, " ")
	}
	fmt.Println("]")

	// Función 1: Definición y llamada
	mostrarTexto()

	// Función 2: Multiplicar 5 y 8 e imprimir el resultado
	multiplicar()

	// Función 3: Multiplicar con variables globales
	aGlobal := 5
	bGlobal := 8
	multiplicarConVariablesGlobales(aGlobal, bGlobal)

	// Función 4: Multiplicar con return y usar el resultado fuera de la función
	resultado := multiplicarConReturn()
	fmt.Println(resultado + 5)

	// Función 5: Validar si el valor de a es 5
	dato := validarDato(aGlobal)
	fmt.Println(dato)
}