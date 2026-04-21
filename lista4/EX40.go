package main

import "fmt"

func main() {
	preco := 6.0
	ingressos := 130.0
	despesa := 300.0

	var maxLucro float64
	var melhorPreco float64
	var melhorQtd float64

	for preco >= 1.0 {
		lucro := preco*ingressos - despesa

		fmt.Printf("Preço: %.2f | Ingressos: %.0f | Lucro: %.2f\n",
			preco, ingressos, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			melhorPreco = preco
			melhorQtd = ingressos
		}

		preco -= 0.60
		ingressos += 30
	}

	fmt.Println("\nLucro máximo:", maxLucro)
	fmt.Println("Melhor preço:", melhorPreco)
	fmt.Println("Ingressos:", melhorQtd)
}
