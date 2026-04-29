package main

import (
	"fmt"
)

const (
	APROBADO   = 6.0
	ASISTENCIA = 70
)

type Estudiante struct {
	Nombre         string
	Calificaciones [5]float64
	Asistencia     int
}

func main() {
	// Declaración con var
	var estudiantes [10]Estudiante

	// Declaración corta :=
	nombres := [10]string{
		"Ana García", "Luis Pérez", "María López", "Carlos Sánchez", "Laura Torres",
		"David Romero", "Sofía Hernández", "Pedro Díaz", "Elena Cruz", "Jorge Martínez",
	}

	// Inicializar datos con for tradicional
	for i := 0; i < len(estudiantes); i++ {
		estudiantes[i].Nombre = nombres[i]
		for j := 0; j < 5; j++ {
			estudiantes[i].Calificaciones[j] = float64((i+j)%10) + 1 // valores simples 1–10
		}
		estudiantes[i].Asistencia = 60 + i*4 // valores entre 60 y 100
	}

	fmt.Println("=== SISTEMA DE CALIFICACIONES ===")

	// Uso de for con range
	for _, e := range estudiantes {
		var suma float64
		for _, nota := range e.Calificaciones {
			suma += nota
		}
		promedio := suma / float64(len(e.Calificaciones))

		// Mostrar calificaciones individuales
		fmt.Printf("%s - Calificaciones: ", e.Nombre)
		for _, nota := range e.Calificaciones {
			fmt.Printf("%.1f ", nota)
		}
		fmt.Printf(" | Promedio: %.2f | Asistencia: %d%%\n", promedio, e.Asistencia)

		// If/else anidados básicos
		if promedio >= APROBADO {
			if e.Asistencia >= ASISTENCIA {
				if promedio >= 9 {
					fmt.Println("Resultado: Excelente")
				} else {
					fmt.Println("Resultado: Aprobado")
				}
			} else {
				fmt.Println("Resultado: Aprobado con baja asistencia")
			}
		} else {
			if e.Asistencia < ASISTENCIA {
				fmt.Println("Resultado: Reprobado por notas y asistencia")
			} else {
				fmt.Println("Resultado: Reprobado")
			}
		}

		// Switch sencillo
		switch {
		case promedio >= 9:
			fmt.Println("Clasificación: A")
		case promedio >= 8:
			fmt.Println("Clasificación: B")
		case promedio >= 7:
			fmt.Println("Clasificación: C")
		case promedio >= 6:
			fmt.Println("Clasificación: D")
		default:
			fmt.Println("Clasificación: F")
		}
		fmt.Println()
	}

	// Estadísticas usando for como while
	var i int
	var sumaGeneral float64
	var mejor float64 = 0
	var peor float64 = 10
	var mejorNombre, peorNombre string
	var aprobados int
	totalNotas := 0

	for i < len(estudiantes) { // estilo while
		var suma float64
		for _, nota := range estudiantes[i].Calificaciones {
			suma += nota
			if nota > mejor {
				mejor = nota
				mejorNombre = estudiantes[i].Nombre
			}
			if nota < peor {
				peor = nota
				peorNombre = estudiantes[i].Nombre
			}
		}
		promedio := suma / 5
		sumaGeneral += suma
		totalNotas += 5
		if promedio >= APROBADO && estudiantes[i].Asistencia >= ASISTENCIA {
			aprobados++
		}
		i++ // incremento manual
	}

	promedioGeneral := sumaGeneral / float64(totalNotas)
	porcentajeAprobados := float64(aprobados) / float64(len(estudiantes)) * 100

	fmt.Println("=== ESTADÍSTICAS ===")
	fmt.Printf("Promedio general: %.2f\n", promedioGeneral)
	fmt.Printf("Mejor calificación: %.1f (%s)\n", mejor, mejorNombre)
	fmt.Printf("Peor calificación: %.1f (%s)\n", peor, peorNombre)
	fmt.Printf("Aprobados: %.1f%%\n", porcentajeAprobados)
}
