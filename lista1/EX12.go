// Uma locadora de charretes cobra R$ 10,00 de taxa para cada 3 horas de uso de uma charrete e R$5,00
// para cada 1 hora abaixo dessas 3 horas. Fazer um programa que leia a quantidade de horas que a charrete
// foi usada e que calcule e escreva quanto o cliente tem de pagar.
// Entrada
// O programa deve ler uma única linha na entrada, contendo o número de horas que o locatário utilizou a
// charrete.
// Saída
// O programa deve imprimir uma linha contendo a frase: O VALOR A PAGAR E = X, onde X é o valor
// do aluguel e deve conter no máximo 2 casas decimais. Após o valor do aluguel o programa deve imprimir
// um caractere de quebra de linha: ‘\n’.

package main

import (
	"fmt"
	"math"
)

// *********** USANDO MATH NA RESOLUÇÃO***********
func main() {
	qtdHoras, taxa3, taxa1 := 0.0, 0.0, 0.0
	valorTotal := 0.0

	fmt.Print("Informe a quantidade de horas: ")
	fmt.Scan(&qtdHoras)

	//Math.Modf - Em Go, a forma mais eficiente e comum de separar a parte inteira da decimal é usando a função Modf do pacote math. Essa função recebe um número do tipo float64 e retorna dois valores separadamente.
	//Math.Trunc - Em go, serve para pegar somente o valor inteiro da variável, jogando o restante após a vírgula fora.
	taxa3 = math.Trunc((qtdHoras / 3))
	taxa1 = qtdHoras - (taxa3 * 3)

	valorTotal = (taxa1 * 5) + (taxa3 * 10)

	fmt.Print(taxa3, taxa1, "\n")

	fmt.Printf("O VALOR A PAGAR E = %.2f!\n", valorTotal)

}

//*********** SEM MATH NA RESOLUÇÃO ***********

// func main() {
// 	qtdHoras, taxaHora3, taxaHora2 := 0, 0, 0
// 	valorCobrado := 0.0

// 	fmt.Print("Informe a quantidade de horas: ")
// 	fmt.Scan(&qtdHoras)

// 	taxaHora3 = qtdHoras / 3

// 	taxaHora2 = qtdHoras % 3

// 	valorCobrado = float64(taxaHora2*5) + float64(taxaHora3*10)

// 	fmt.Printf("O VALOR COBRADO E %.2f!\n", valorCobrado)

// }
