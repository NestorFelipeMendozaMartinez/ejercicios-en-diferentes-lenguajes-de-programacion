package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Definición de las variables
	a := 10
	b := 4
	var c int

	// Multiplicación
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: *")
	fmt.Println("La segunda variable es:", b)
	c = a * b
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))

	// División
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: /")
	fmt.Println("La segunda variable es:", b)
	c = a / b
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))

	// Suma
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: +")
	fmt.Println("La segunda variable es:", b)
	c = a + b
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))

	// Resta
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: -")
	fmt.Println("La segunda variable es:", b)
	c = a - b
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))

	// Potencia
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: **")
	fmt.Println("La segunda variable es:", b)
	c = int(math.Pow(float64(a), float64(b))) // Uso de math.Pow() para calcular la potencia
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))

	// Módulo
	fmt.Println("La primera variable es:", a)
	fmt.Println("El signo de la operación es: %")
	fmt.Println("La segunda variable es:", b)
	c = a % b
	fmt.Println("El resultado es:", c)
	fmt.Println("El tipo de c es:", reflect.TypeOf(c))
}