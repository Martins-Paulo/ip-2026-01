package main

// 6. Crie um programa para determinar se um número inteiro A é divisível por outro número B. Os
// valores devem ser fornecidos pelo usuário.

import "fmt"

func main() {
	numero1, numero2 := 0, 0

	fmt.Print("Informe um número A e número B: ")
	fmt.Scan(&numero1, &numero2)

	if numero1%numero2 == 0 {
		fmt.Printf("O número %d é divisível por %d!\n", numero1, numero2)
	} else {
		fmt.Printf("O número %d não é divisível por %d!\n", numero1, numero2)
	}

}
