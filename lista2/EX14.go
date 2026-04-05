// 14. Desenvolva um programa para calcular e imprimir o preço final de um carro. O valor do preço inicial de
// fábrica é fornecido pelo usuário. O carro pode ter as seguintes opções:
// a) Ar condicionado (R$ 1750,00)
// b) Pintura metálica (R$ 800,00)
// c) Vidro elétrico (R$ 1200,00)
// d) Direção hidráulica (R$ 2000,00)

package main

import (
	"fmt"
	"time"
)

func main() {

	var opcao int
	arCondicionado, pintura, vidro, direcao := 1750.0, 800.0, 1200.0, 2000.0
	valorTotal, valorInicial := 0.0, 0.0

	fmt.Println("Informe o valor inicial do carro:")
	fmt.Scan(&valorInicial)

	valorTotal = valorInicial

	for {
		fmt.Println("\n------ MENU ACESSÓRIOS ------\n")
		fmt.Println("Escolha opção que queira adicionar: \n1- Ar condicionado\n2- Pintura\n3- Vidro elétrico\n4- Direção hidráulica\n0 - Para finalizar")

		fmt.Scan(&opcao)

		if opcao == 0 {
			break
		}

		switch opcao {
		case 1:
			valorTotal += arCondicionado
		case 2:
			valorTotal += pintura
		case 3:
			valorTotal += vidro
		case 4:
			valorTotal += direcao
		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Printf("Subtotal: R$ %.2f\n", valorTotal)
		time.Sleep(2 * time.Second)

	}

	fmt.Printf("O valor total é R$ %.2f!\n", valorTotal)

}
