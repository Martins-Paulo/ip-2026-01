// Faça um programa que leia um valor n, inteiro e positivo, calcule e mostre a seguinte somaFinal:

// S =Xn,k=1=1 k= 1 + 1/2 + 1/3 + 1/4 + . . . + 1/n (3)

// Entrada
// O programa deve ler um número inteiro positivo e maior que 1.
// Saída
// O programa deve apresentar uma linha contendo o valor final do somatorio com 6 casas decimais. Caso
// o número lido não atenda as especificações da entrada, o programa deve apresentar a mensagem: "Numero
// invalido!".

package main

import (
	"fmt"
)

func main() {
	var numero int
	var somaFinal, somatorio float64

	fmt.Print("Digite um número: \n")
	fmt.Scan(&numero)

	if numero > 0 {
		for i := 1; i <= numero; i++ {

			somatorio = 1.0 / float64(i)
			somaFinal += somatorio

		}
		fmt.Printf("%.6f", somaFinal)

	} else {
		fmt.Print("Número inválido!\n")
	}

}
