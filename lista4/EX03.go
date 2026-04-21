package main

import "fmt"

func main() {
	var salarioCarlos float64
	fmt.Print("Digite o salário do Carlos: ")
	fmt.Scan(&salarioCarlos)

	salarioJoao := salarioCarlos / 3
	meses := 0

	// Enquanto o valor de João for menor que o de Carlos
	for salarioJoao < salarioCarlos {
		salarioCarlos += salarioCarlos * 0.02 // Rendimento de 2%
		salarioJoao += salarioJoao * 0.05     // Rendimento de 5%
		meses++
	}

	fmt.Printf("Serão necessários %d meses para João igualar ou ultrapassar Carlos.\n", meses)
	fmt.Printf("Valor final Carlos: R$ %.2f\n", salarioCarlos)
	fmt.Printf("Valor final João: R$ %.2f\n", salarioJoao)
}
