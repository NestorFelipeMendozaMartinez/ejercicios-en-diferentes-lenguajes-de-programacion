package main

import "fmt"

func main() {
    // Primer bloque: Contador con while (simulado con for)
    contador := 0
    for contador < 30 {
        contador++
        fmt.Println("Iteración", contador)
    }

    // Segundo bloque: Combinado con if-else
    var a int
    for {
        fmt.Println("Introduzca un valor mayor a 10:")
        fmt.Scanln(&a)

        if a > 10 {
            fmt.Println("Es correcto")
        } else {
            fmt.Println("Es incorrecto, pruebe de nuevo")
            break // Sale del ciclo for si el valor ingresado es incorrecto
        }
    }
}