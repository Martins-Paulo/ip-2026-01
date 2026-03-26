package main

import "fmt"

//2. Crie um programa que leia um valor inteiro e diga se ele é positivo, negativo ou nulo.

func main() {
	numero := 0

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Printf("O número %d é positivo!\n", numero)
	} else if numero < 0 {
		fmt.Printf("O número %d é negativo!\n", numero)
	} else {
		fmt.Printf("O número %d é nulo!\n", numero)
	}

}
