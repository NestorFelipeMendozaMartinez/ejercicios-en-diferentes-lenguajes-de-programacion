package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Función para obtener el ganador del partido
func ganadoresOctavos(marcador string, equipo1 string, equipo2 string) string {
	// Dividir el marcador en goles
	parts := strings.Split(marcador, "-")
	goles1, _ := strconv.Atoi(parts[0])
	goles2, _ := strconv.Atoi(parts[1])

	// Determinar el ganador
	if goles1 > goles2 {
		return equipo1
	} else {
		return equipo2
	}
}

func main() {
	// Lista de partidos
	equipos := []string{
		"1. Países_Bajos vs. Estados_Unidos",
		"2. Argentina vs. Australia",
		"3. Francia vs. Polonia",
		"4. Inglaterra vs. Senegal",
		"5. Japón vs. Croacia",
		"6. Brasil vs. Corea_del_Sur",
		"7. Marruecos vs. España",
		"8. Portugal vs. Suiza",
	}

	// Mapa de resultados
	resultados := make(map[string]string)

	// Imprimir el encabezado
	fmt.Println("Mundial Qatar 2022\n\nOctavos de final\n")

	// Pedir marcadores y almacenar resultados
	for _, partido := range equipos {
		fmt.Println(partido)
		fmt.Print("Ingrese el marcador final (Formato: 2-1): ")
		var marcador string
		fmt.Scanln(&marcador)
		resultados[partido] = marcador
	}

	// Imprimir los ganadores
	fmt.Println("\nGanadores de la fase de octavos:")
	for partido, marcador := range resultados {
		// Dividir la cadena para extraer los equipos
		parts := strings.Split(partido, " vs. ")
		equipo1 := parts[0][3:] // Obtener el primer equipo (quitando el número)
		equipo2 := parts[1]     // Obtener el segundo equipo

		// Obtener el ganador
		ganador := ganadoresOctavos(marcador, equipo1, equipo2)
		fmt.Printf("Ganador de %s es: %s\n", partido, ganador)
	}
}