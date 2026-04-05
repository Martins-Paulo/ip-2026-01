// 12. Construa um programa que receba a idade de uma pessoa e classifique-a seguindo o critério apresentado a
// seguir. Considere que a idade é um valor inteiro e que será informada de forma válida.
// Idade Classificação
// 0 a 2 anos Recém-nascido
// 3 a 11 anos Criança
// 12 a 19 anos Adolescente
// 20 a 55 anos Adulto
// Acima de 55 anos Idoso

package main

import "fmt"

func main() {
	idade := 0

	fmt.Println("Informe a idade!")
	fmt.Scan(&idade)

	switch {
	case idade < 0:
		fmt.Println("Idade inválida!")
	case idade <= 2:
		fmt.Println("Recém-nascido")
	case idade <= 11:
		fmt.Println("Criança")
	case idade <= 19:
		fmt.Println("Adolescente")
	case idade <= 55:
		fmt.Println("Adulto")
	default:
		fmt.Println("Idoso")
	}
}
