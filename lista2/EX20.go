// 20. Elabore um programa que calcule o valor a ser pago por um produto considerando o preço normal de etiqueta
// e a escolha da condição de pagamento. Utilize os códigos da tabela a seguir para saber qual a condição de
// pagamento escolhida e efetuar o cálculo adequado.

// Código Condição de Pagamento
// 1 À vista, dinheiro ou cheque, 10% de desconto
// 2 À vista, cartão de crédito, 5% de desconto
// 3 Em 2 vezes, preço normal de etiqueta sem juros
// 4 Em 3 vezes, preço normal de etiqueta + 10% de juros

package main

import (
	"fmt"
)

func main() {
	var opcao int
	var precoBase, precoFinal float64

	fmt.Println("Informe o valor do produto na etiqueta:")
	fmt.Scan(&precoBase)

	fmt.Println("Informe a opção de pagamento:\n1 - À vista, dinheiro ou cheque.\n2 - À vista, cartão de crédito\n3 - Parcelado 2 vezes\n4 - Parcelado 4 vezes.")
	fmt.Scan(&opcao)

	precoFinal = precoBase

	switch opcao {
	case 1:
		precoFinal *= 0.9
	case 2:
		precoFinal *= 0.95
	case 3:
		precoFinal = precoBase
	case 4:
		precoFinal *= 1.1
	default:
		fmt.Println("Opção inválida!")

	}
	fmt.Printf("O valor do pagamento total de acordo com a opção %d é R$ %.2f", opcao, precoFinal)

}
