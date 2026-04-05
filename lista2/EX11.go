// Escreva um programa que receba o valor de x, calcule e imprima o valor de f(x), dado que:

// f(x) = 8 / (2-x)

package main

import "fmt"

func main() {
	numero := 0.0

	fmt.Println("Informe o valor de x: ")
	fmt.Scan(&numero)

	sucesso, resultado := calculo(numero)

	if sucesso {
		fmt.Printf("O f(%.2f) é %.3f!\n", numero, resultado)
	} else {
		fmt.Println("Divisão não permitida")
	}

}

func calculo(x float64) (bool, float64) {
	valido := true
	conta := 0.0
	if x == 2 {
		valido = false
	} else {
		conta = 8 / (2 - x)
	}

	return valido, conta
}
