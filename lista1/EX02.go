package main

import f "fmt"

func main() {
	//Declaração de variavéis
	casosTestes, totalIngressos := 0, 0.0
	const (
		valorPopular, valorGeral, valorArquibancada, valorCadeiras = 1.0, 5.0, 10.0, 20.0
	)

	porcentagemPopular, porcentagemGeral, porcentagemArquibancada, porcentagemCadeiras := 0.0, 0.0, 0.0, 0.0

	qtdPopular, qtdGeral, qtdArquibancada, qtdCadeiras := 0.0, 0.0, 0.0, 0.0
	rendaJogo := 0.0
	// Pegando dados sobre quantidade de dados solicitados.

	f.Print("Informe a quantidade de casos:")
	f.Scan(&casosTestes)

	//Usando for para captar os dados dependendo da quantidade de casos que o usuário informar.
	for i := 1; i <= casosTestes+1; i++ {

		f.Printf("Informe o número total de ingressos para o caso número %d!"+"\nEm seguida informe as porcentagens de ingressos, em sequência, de cada categoria:\nPopular, Geral, Arquibancada e Cadeiras: ", i)
		f.Scan(&totalIngressos, &porcentagemPopular, &porcentagemGeral, &porcentagemArquibancada, &porcentagemCadeiras)

		qtdPopular = totalIngressos * (porcentagemPopular / 100.0)
		qtdGeral = totalIngressos * (porcentagemGeral / 100.0)
		qtdArquibancada = totalIngressos * (porcentagemArquibancada / 100.0)
		qtdCadeiras = totalIngressos * (porcentagemCadeiras / 100.0)

		rendaJogo = (qtdPopular * valorPopular) + (qtdGeral * valorGeral) + (qtdArquibancada * valorArquibancada) + (qtdCadeiras * valorCadeiras)

		f.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, rendaJogo)
	}

}
