package main

import "fmt"

func main() {
    // Crear la lista de nombres
    nombres := []string{"Esteban", "Hans", "Jhon", "Juan Pablo"}

    // Imprimir los nombres
    for _, nombre := range nombres {
        fmt.Println(nombre)
    }

    // Crear la lista de personas (como slice de mapas)
    personas := []map[string]string{
        {"nombre": "Esteban", "Edad": "28"},
        {"nombre": "Hans", "Edad": "27"},
        {"nombre": "Jhon", "Edad": "41"},
        {"nombre": "Juan Pablo", "Edad": "23"},
    }

    // Imprimir los nombres y edades
    for _, persona := range personas {
        fmt.Println(persona["nombre"], persona["Edad"])
    }
}