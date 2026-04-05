// Escrever um programa que leia o número de identificação, as 3 notas obtidas por um aluno nas 3 verificações e
// a média dos exercícios que fazem parte da avaliação. Calcular a média de aproveitamento do aluno, usando a
// fórmula:

// O programa deve escrever o número do aluno, suas notas, a média dos exercícios, a média de aproveitamento,
// o conceito correspondente e a mensagem: APROVADO se o conceito for A, B ou C e REPROVADO, se o
// conceito for D ou E.

package main

import (
	"fmt"
)

func main() {
	var nota1, nota2, nota3, mediaExercicios, mediaFinal float64
	var conceito, aprovacao string
	var idAluno int

	fmt.Println("Informe o número de identificação do Aluno(a):")
	fmt.Scan(&idAluno)

	fmt.Println("Informe a nota 1, nota 2 e nota 3 do aluno(a):")
	fmt.Scan(&nota1, &nota2, &nota3)

	fmt.Println("Informe a média dos exercícios do aluno")
	fmt.Scan(&mediaExercicios)

	mediaFinal = (nota1 + (nota2 * 2) + (nota3 * 3) + mediaExercicios) / 7

	if mediaFinal >= 6 {
		aprovacao = "APROVADO"
	} else {
		aprovacao = "REPROVADO"
	}

	switch {
	case mediaFinal < 4:
		conceito = "E"
	case mediaFinal < 6:
		conceito = "D"
	case mediaFinal < 7.5:
		conceito = "C"
	case mediaFinal < 9:
		conceito = "B"
	case mediaFinal <= 10:
		conceito = "A"
	default:
		fmt.Println("Nota inválida!")
	}

	fmt.Printf("Identificação do Aluno(a): %d\n", idAluno)
	fmt.Printf("Notas:\nNota 1: %.2f Nota 2: %.2f Nota 3: %.2f\n", nota1, nota2, nota3)
	fmt.Printf("Média Exercícios: %.2f\n", mediaExercicios)
	fmt.Printf("Média de Aproveitamento: %.2f\n", mediaFinal)
	fmt.Println("Conceito: ", conceito)
	fmt.Println("Status: ", aprovacao)

}
