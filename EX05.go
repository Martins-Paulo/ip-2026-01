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
	fmt.Print(tipoCliente)

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
