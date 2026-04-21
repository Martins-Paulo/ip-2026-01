package main

import (
	"fmt"
)

func main() {
	var a1, a2, n int

	fmt.Print("Digite o primeiro termo: ")
	fmt.Scan(&a1)

	fmt.Print("Digite o segundo termo: ")
	fmt.Scan(&a2)

	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("N deve ser maior ou igual a 3.")
		return
	}

	fmt.Printf("Série de Fetuccine: %d %d ", a1, a2)

	// termos iniciais
	prev2 := a1
	prev1 := a2

	for i := 3; i <= n; i++ {
		var atual int

		if i%2 == 1 {
			// ímpar: soma
			atual = prev1 + prev2
		} else {
			// par: subtração
			atual = prev1 - prev2
		}

		fmt.Printf("%d ", atual)

		// atualiza os anteriores
		prev2 = prev1
		prev1 = atual
	}
}
