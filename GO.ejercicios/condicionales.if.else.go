package main

import "fmt"

func main() {
    // Primer bloque: Verificar si 'a' es igual a 2
    a := 2
    if a == 2 {
        fmt.Println("a vale 2")
    } else {
        fmt.Println("a es diferente de 2")
    }

    // Segundo bloque: Concatenar condiciones con '&&' (equivalente a 'and' en C++) y '!=' (equivalente a 'not' en C++)
    nombre := "Esteban"
    edad := 28
    pais := "Colombia"

    if nombre == "Esteban" && edad == 28 && pais == "Colombia" {
        fmt.Printf("Su nombre es %s, tiene %d y es de %s\n", nombre, edad, pais)
    } else if nombre == "Esteban" && edad != 28 {
        fmt.Println("Su nombre es Esteban y no tiene 28 años")
    } else if nombre != "Esteban" && edad == 28 {
        fmt.Println("Su nombre no es Esteban y tiene 28 años")
    } else {
        fmt.Println("No se llama Esteban ni tiene 28 años")
    }
}