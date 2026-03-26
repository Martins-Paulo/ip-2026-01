// Desenvolver um programa que leia um número inteiro e verifique se o número é divisível por três e
// também é divisível por cinco.
// Entrada
// O programa deve ler uma linha contendo um número inteiro na entrada.
// Saída
// O programa deve imprimir uma linha contendo a frase: O NUMERO E DIVISIVEL, se ele for divisível
// tanto por três quanto por cinco, ou a frase O NUMERO NAO E DIVISIVEL, em caso contrário. Após
// imprimir uma das frases, o programa deve imprimir um caractere de quebra de linha: ‘\n’.
package main

import "fmt"

func main() {
	numero := 0

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Printf("O número %d é divisível por 3 e 5!\n", numero)
	} else {
		fmt.Printf("O número %d não é divisível por 3 e 5!\n", numero)
	}
}
