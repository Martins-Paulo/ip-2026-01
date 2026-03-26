package main

// 5. Escreva um programa que leia um número inteiro e diga se ele é ou não um número divisível por 5.

import "fmt"

func main() {
	numero := 0

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if numero%5 == 0 {
		fmt.Printf("O número %d é divisível por 5!\n", numero)
	} else {
		fmt.Printf("O número %d não é divisível por 5!\n", numero)
	}
}
