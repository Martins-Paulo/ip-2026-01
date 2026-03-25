package main

import "fmt"

func main() {
	var valorA, valorB, valorC, valorD, determinante float64

	fmt.Print("Informe os valores de A, B, C e D: ")
	fmt.Scan(&valorA, &valorB, &valorC, &valorD)

	determinante = (valorA * valorD) - (valorC * valorB)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)

}
