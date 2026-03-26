package main

//1. Escreva um programa que leia um valor inteiro e diga se o número informado é par ou ímpar.

import "fmt"

func main() {
	numero := 0

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Printf("O número %d é Par!\n", numero)
	} else {
		fmt.Printf("O número %d é Impar!\n", numero)
	}
}
