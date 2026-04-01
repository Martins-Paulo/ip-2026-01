// Faça um programa que indique se um número inteiro informado pelo usuário está compreendido
// entre 20 e 90 ou não. (20 e 90 não estão na faixa de valores).

package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Informe um número: ")
	_, err := fmt.Scan(&numero)

	if err != nil {
		fmt.Println("Valor inválido\nInforme apenas números!\n")
		return
	}

	if numero > 20 && numero < 90 {
		fmt.Printf("O numero %d está compreendido entre 20 e 90!\n", numero)
	} else {
		fmt.Printf("O numero %d não está compreendido entre 20 e 90!\n", numero)
	}

}
