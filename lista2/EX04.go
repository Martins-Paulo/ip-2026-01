package main

// 4. Faça um programa que leia um número do tipo float e imprima sua raiz quadrada caso o mesmo seja
// positivo ou nulo. Caso ele seja negativo, mostre o seu quadrado.

import (
	"fmt"
	"math"
)

func main() {
	numero := 0.0

	fmt.Print("Informe o número: ")
	fmt.Scan(&numero)

	if numero >= 0 {
		numero = math.Sqrt(numero)
		fmt.Printf("A raiz quadrada do número é %.2f\n", numero)
	} else {
		numero = math.Pow(numero, 2)
		fmt.Printf("O quadrado do número é %.2f!\n", numero)
	}

}
