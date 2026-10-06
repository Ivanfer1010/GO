// Ivan Gualotuña

package main

import (
	"fmt"
)

// Ejercicio 1 · Array de números primos
func ejercicio1() {
	primos := [5]int{2, 3, 5, 7, 11}
	primero := primos[0]
	centro := primos[len(primos)/2]
	ultimo := primos[len(primos)-1]
	cantidad := len(primos)

	fmt.Printf("1. Primero: %d Centro: %d Último: %d Cantidad: %d\n", primero, centro, ultimo, cantidad)
}

// Ejercicio 2 · Array de muestras y promedio
func ejercicio2() {
	muestras := [3]float64{71.8, 56.2, 89.5}
	suma := 0.0

	for _, valor := range muestras {
		suma += valor
	}

	promedio := suma / float64(len(muestras))
	fmt.Printf("2. Promedio: %.2f\n", promedio)
}

// Ejercicio 3 · Slices y cortes
func ejercicio3() {
	primos := []int{2, 3, 5}
	primos = append(primos, 7, 11)

	// Nota: En Go, el índice final del corte es exclusivo ([ini:fin]), 
	// por lo que [1:4] incluye los índices 1, 2 y 3 para obtener [3 5 7].
	corte := primos[1:4]

	fmt.Printf("3. Primos: %v corte: %v\n", primos, corte)
}

// Ejercicio 4 · Maps para conteo de votos
func ejercicio4() {
	votosRecibidos := []string{"Amber", "Brian", "Amber", "Brian", "Amber"}
	votos := make(map[string]int)

	for _, candidato := range votosRecibidos {
		votos[candidato]++
	}

	fmt.Printf("4. Votos: %v\n", votos)
}

func main() {
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
}