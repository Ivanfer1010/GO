// Ivan Gualotuña

package main

import (
	"fmt"
)

// Ejercicio 1 · Array de números primos
func ejercicioArrayPrimos() {
	primos := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Primer valor: ", primos [0], "Centro: ", primos [2], "Ultimo: ", primos[4], "Cantidad", len(primos))
}
// Ejercicio 2 · Array de muestras y promedio
func ejercicioArrayPromedio() {
	muestras := [3]float64{71.8, 56.2, 89.5}
	suma := 0.0
	for _, m := range muestras {
		suma += m
	}
	fmt.Printf("Promedio: %.2f\n", suma/float64(len(muestras)) )
}
// Ejercicio 3 · Slices y cortes
func ejercicioArrayPrimoSlice() {
	primos := []int{2, 3, 5}
	primos = append(primos, 7, 11)
	fmt.Println("Primos", primos, "Corte", primos [1:4])
}
func ejercicioMap() {
	votos := []string{"Amber", "Brian", "Amber", "Brian", "Amber"}
	conteo := make(map[string]int)
	for _, v := range votos {
		conteo[v]++
	}

	fmt.Println("Votos", conteo)
}
func main() {
	ejercicioArrayPrimos()
	ejercicioArrayPromedio()
	ejercicioArrayPrimoSlice()
	ejercicioMap()
	}
