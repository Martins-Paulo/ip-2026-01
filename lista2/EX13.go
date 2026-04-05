// 13. Escreva um programa que receba um número inteiro positivo de 3 casas e imprima o algarismo da casa das
// dezenas. Não se esqueça de testar para ver se o número informado tem realmente 3 casas.

package main

import "fmt"

func main() {
	numero := 0
	//dezena := 0
	//centena := 0

	fmt.Println("Informe um número:")
	fmt.Scan(&numero)

	// if numero < 100 || numero > 999 {
	// 	fmt.Println("Número inválido!")
	// } else {
	// 	centena = numero / 100
	// 	dezena = (numero - (centena * 100)) / 10

	// 	fmt.Printf("O algarismo da casa das dezenas do número %d é %d\n", numero, dezena)
	// }
	sucesso, dezena := extrairdezenas(numero)

	if sucesso {
		fmt.Printf("O algarismo da casa das dezenas do número %d é %d\n", numero, dezena)
	} else {
		fmt.Println("Número inválido!")
	}

}

func extrairdezenas(numero int) (bool, int) {
	dezena := 0

	if numero < 100 || numero > 999 {
		return false, 0
	} else {
		dezena = (numero / 10) % 10

	}
	return true, dezena
}
