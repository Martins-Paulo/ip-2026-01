// Escreva um programa que leia o destino de um passageiro e se a viagem inclui retorno (ou seja, se é
// só ida, ou se é ida e volta). Informe o preço da passagem de acordo com a tabela a seguir.
// Destino Ida Ida e Volta
// 1 - Região Norte R$ 500,00 R$ 900,00
// 2 - Região Nordeste R$ 350,00 R$ 650,00
// 3 - Região Centro-Oeste R$350,00 R$ 600,00
// 4 - Região Sul R$ 300,00 R$ 550,00

// Sugestão: Considere “1” representando a Região Norte, “2” para Nordeste, “3” para Centro-Oeste e
// “4” para Sul. Para ver se a viagem inclui retorno, considere “1”- para sim (ida e volta) e “2” - para
// não (só ida).
// Cuidado com valores inválidos tanto para o destino quanto para o fato de incluir ou não o retorno.

package main

import "fmt"

func main() {
	var destino, idaVolta int
	var preco float64

	fmt.Println("Informe o Destino do passageiro:\n1 - Norte\n2 - Nordeste\n3 - Centro-Oeste\n4 - Sul\nEm seguida informe:\n1 -  Ida e volta\n2 - Somente ida: \n")
	_, err := fmt.Scan(&destino, &idaVolta)
	//Verifica se foi digitado um palavra e retorna mensagem de erro.
	if err != nil {
		fmt.Println("Por favor informe um número válido!\n(1 ao 5 para Destino)\n1 - para sim ida e volta\n2 - para ida e volta\n")
		return
	}
	//Verifica se foi digitado um valor que nao tenha nenhuma regiao
	if destino < 1 || destino > 4 {
		fmt.Println("Valor inválido para destino!\nPor favor escolha entre os números 1, 2, 3 ou 4!\n")
		return
	}
	if idaVolta < 1 || idaVolta > 2 {
		fmt.Println("Valor inválido para tipo de viagem!\nPor favor escolha entre os números 1 ou 2\n")
		return
	}
	//Usando if e else para fazer a logica. Fica muito longo
	if destino == 1 {
		if idaVolta == 2 {
			fmt.Println("O valor somente de ida é R$ 500,00 reais\n")
		} else {
			fmt.Println("O valor de ida e volta é R$ 900,00 reais\n")
		}
	} else if destino == 2 {
		if idaVolta == 2 {
			fmt.Println("O valor somente de ida é R$ 350,00 reais\n")
		} else {
			fmt.Println("O valor de ida e volta é R$ 650,00 reais\n")
		}
	} else if destino == 3 {
		if idaVolta == 2 {
			fmt.Println("O valor somente de ida é R$ 350,00 reais\n")
		} else {
			fmt.Println("O valor de ida e volta é R$ 600,00 reais\n")
		}
	} else {
		if idaVolta == 2 {
			fmt.Println("O valor somente de ida é R$ 300,00 reais\n")
		} else {
			fmt.Println("O valor de ida e volta é R$ 550,00 reais\n")
		}
	}
	//Usando o switch case:
	switch destino {
	case 1:
		if idaVolta == 1 {
			preco = 900.00
		} else {
			preco = 500.00
		}
	case 2:
		if idaVolta == 1 {
			preco = 650.00
		} else {
			preco = 350.00
		}
	case 3:
		if idaVolta == 1 {
			preco = 600.00
		} else {
			preco = 350.00
		}
	case 4:
		if idaVolta == 1 {
			preco = 550.00
		} else {
			preco = 300.00
		}
	}
	fmt.Printf("O valor da passagem é R$ %.2f\n", preco)

}
