package main

import "fmt"

func main() {
    // Crear un "diccionario" utilizando map en Go
    coche := make(map[string]string)

    // Agregar elementos al map
    coche["marca"] = "Ford"
    coche["color"] = "rojo"
    coche["modelo"] = "sedan"
    coche["placa"] = "ROS227"

    // Mostrar el valor de 'color'
    fmt.Println(coche["color"])

    // Cambiar el valor de 'color'
    coche["color"] = "verde"
    fmt.Println(coche["color"])

    // Mostrar el valor de 'marca'
    fmt.Println(coche["marca"])

    // Cambiar el valor de 'marca'
    coche["marca"] = "Renault"
    fmt.Println(coche["marca"])

    // Mostrar todo el map
    for clave, valor := range coche {
        fmt.Println(clave, ":", valor)
    }
}