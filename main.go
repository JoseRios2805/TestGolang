package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deliveryBikes(N int, packageWeights []int) int {

	sortedWeights := make([]int, len(packageWeights))
	copy(sortedWeights, packageWeights)
	sort.Ints(sortedWeights)

	numBikes := 0
	currentWeight := 0
	for _, weight := range sortedWeights {
		if currentWeight+weight <= N {
			currentWeight += weight
		} else {
			numBikes++
			currentWeight = weight
		}
	}
	numBikes++ //Añade una bicicleta más para el último paquete.

	return numBikes
}

func main() {
	N := 10
	packageWeights := []int{5, 3, 6, 8, 2}
	fmt.Printf("Numero de bicicletas requeridas: %d\n", deliveryBikes(N, packageWeights))
}