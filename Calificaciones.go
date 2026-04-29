package main

import (
	"fmt"
)

func main() {
	// Variables
	var numEstudiantes int
	fmt.Print("¿Cuántos estudiantes deseas registrar? ")
	fmt.Scanln(&numEstudiantes)

	// Arrays para nombres y calificaciones
	nombres := make([]string, numEstudiantes)
	calificaciones := make([][]float64, numEstudiantes)

	// Ingreso de datos
	for i := 0; i < numEstudiantes; i++ {
		fmt.Printf("\nNombre del estudiante %d: ", i+1)
		fmt.Scanln(&nombres[i])

		var numNotas int
		fmt.Printf("¿Cuántas calificaciones tiene %s? ", nombres[i])
		fmt.Scanln(&numNotas)

		calificaciones[i] = make([]float64, numNotas)
		for j := 0; j < numNotas; j++ {
			fmt.Printf("Ingresa la calificación %d de %s: ", j+1, nombres[i])
			fmt.Scanln(&calificaciones[i][j])
		}
	}

	// Procesamiento de datos
	var sumaGeneral float64
	var totalNotas int

	for i := 0; i < numEstudiantes; i++ {
		var suma float64
		for _, nota := range calificaciones[i] {
			suma += nota
		}
		promedio := suma / float64(len(calificaciones[i]))
		fmt.Printf("\nPromedio de %s: %.2f\n", nombres[i], promedio)

		// If/Else para desempeño
		if promedio >= 90 {
			fmt.Println("Desempeño: Excelente")
		} else if promedio >= 70 {
			fmt.Println("Desempeño: Bueno")
		} else {
			fmt.Println("Desempeño: Necesita mejorar")
		}

		// Switch para clasificación
		switch {
		case promedio >= 90:
			fmt.Println("Clasificación: A")
		case promedio >= 80:
			fmt.Println("Clasificación: B")
		case promedio >= 70:
			fmt.Println("Clasificación: C")
		case promedio >= 60:
			fmt.Println("Clasificación: D")
		default:
			fmt.Println("Clasificación: F")
		}

		// Acumular para promedio general
		sumaGeneral += suma
		totalNotas += len(calificaciones[i])
	}

	// Promedio general de todos los estudiantes
	promedioGeneral := sumaGeneral / float64(totalNotas)
	fmt.Printf("\nPromedio general del grupo: %.2f\n", promedioGeneral)
}
