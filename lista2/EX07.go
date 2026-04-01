// Escreva um programa que leia três valores inteiros distintos (assuma que o usuário digitará valores
// diferentes entre si) e os armazene nas variáveis A, B e C. Em seguida, descubra o menor valor e o
// armazene em uma variável denominada MENOR; o maior valor, coloque-o na variável MAIOR e o
// valor intermediário, na variável INTER. Imprima os valores em ordem crescente, ou seja, imprima
// as variáveis MENOR, INTER e MAIOR, nessa ordem.

package main

import "fmt"

func main() {
	var A, B, C, maior, inter, menor int

	fmt.Print("Informe os Valores para A, B e C, em sequência: \n")
	fmt.Scan(&A, &B, &C)
	//Verificação do maior número
	if A > B && A > C {
		maior = A
	} else if B > C {
		maior = B
	} else {
		maior = C
	}
	//Verificação do menor número
	if A < B && A < C {
		menor = A
	} else if B < C {
		menor = B
	} else {
		menor = C
	}
	//Verificação do inter, pensei em 2 metodos:
	//1 - Usando if e else
	// if A != maior && A != menor {
	// 	inter = A
	// } else if B != maior && B != menor {
	// 	inter = B
	// } else {
	// 	inter = C
	// }
	//Usando Soma
	inter = (A + B + C) - (menor + maior)

	fmt.Printf("Valores em ordem Crescente: %d %d %d\n", menor, inter, maior)

}
