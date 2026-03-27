// Fazer um algoritmo que calcule e imprima o salário reajustado de um funcionário de acordo com as
// seguintes regras:
// • Salário de até R$ 300,00, reajuste de 50%;
// • Salário maior que R$300,00 reajuste de 30%;
// Entrada
// A entrada conterá uma linha com o salário do funcionário.
// Saída
// A saída deve conter, numa linha com a frase: SALARIO COM REAJUSTE = x, onde x é um valor real
// com duas casas decimais e corresponde ao valor do salário reajustado. Logo em seguida ao valor de x, o
// programa devem imprimir o caractere de quebra de linha: ‘\n’.

package main

import "fmt"

func main() {
	var salario, salarioReajustado float64

	fmt.Print("Informe o salário do funcionário: ")
	// "_" é como se fosse um local de memoria que joga imediatamente o que foi adicionado nele fora.
	// nesse caso a função Scan sempre pega 2 valores, a quantidade(valor digitado) e erro(se deu certo ou nop)
	_, err := fmt.Scan(&salario)

	fmt.Print(err)

	if err != nil {
		fmt.Print("Erro! Por favor digite um número!\n")
	} else {
		switch {
		case salario <= 300:
			salarioReajustado = salario + (salario * 0.5)
			fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
		case salario > 300:
			salarioReajustado = salario + (salario * 0.3)
			fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)

		}

	}

}
