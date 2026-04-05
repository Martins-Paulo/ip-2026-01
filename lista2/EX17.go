// 17. Desenvolver um programa para calcular a conta de água para a SANEAGO. O custo da água varia dependendo
// do tipo do consumidor - residencial, comercial ou industrial. A regra para calcular a conta é:
// • Residencial: R$ 5,00 de taxa mais R$ 0,05 por m3 gastos;
// • Comercial: R$ 500,00 para os primeiros 80 m3 gastos mais R$ 0,25 por m3 gastos acima dos 80 m3;
// • Industrial: R$ 800,00 para os primeiros 100 m3 gastos mas R$ 0,04 por m3 gastos acima dos 100 m3;
// O programa deverá ler a conta do cliente, seu tipo (residencial, comercial e industrial) e o seu consumo de água em
// metros cúbicos. Como resultado imprimir a conta do cliente e o valor em real a ser pago pelo mesmo.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var contaCliente int64
	var consumoAgua, valorConta float64
	//Rune é usado no lugar de char, tem maior capacidade e pode colocar emogi, simbolos.
	var tipoCliente string

	fmt.Print("Informe o número da conta, valor do consumo de água e tipo de cliente: 'C' - Comercial, 'I' - Industrial ou 'R' - Residencial. ")
	fmt.Scan(&contaCliente, &consumoAgua, &tipoCliente)

	tipoCliente = strings.ToUpper(tipoCliente)

	if tipoCliente == "I" {
		valorConta = 800 + ((consumoAgua - 100) * 0.04)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", contaCliente, valorConta)
	} else if tipoCliente == "C" {
		valorConta = 500 + ((consumoAgua - 80) * 0.25)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", contaCliente, valorConta)
	} else if tipoCliente == "R" {
		valorConta = 5 + (consumoAgua * 0.05)
		fmt.Printf("Conta = %d\nVALOR DA CONTA = %.2f", contaCliente, valorConta)
	} else {
		fmt.Print("Informe um tipo de cliente válido!")
	}
}
