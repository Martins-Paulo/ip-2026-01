package main

import (
	"fmt"
	"math"
)

func main() {
	var valorA, valorB, valorC, delta float64

	fmt.Print("Informe os valores de A, B e C: ")
	fmt.Scan(&valorA, &valorB, &valorC)

	delta = (math.Pow(valorB, 2)) - (4 * valorA * valorC)

	fmt.Printf("O VALOR DE DELTA E = %.2f", delta)

}
