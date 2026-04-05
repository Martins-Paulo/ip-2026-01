// 15. Faça um programa que leia uma data (dia, mês e ano, em uma variável inteira cada), e escreva a mesma data
// no formato dia de (mês por extenso) de ano.

package main

import "fmt"

func main() {
	dia, mes, ano := 0, 0, 0
	var mesExtenso string

	fmt.Println("Informe o dia, mês e ano:")
	fmt.Scan(&dia, &mes, &ano)

	switch mes {
	case 1:
		mesExtenso = "Janeiro"
	case 2:
		mesExtenso = "Fevereiro"
	case 3:
		mesExtenso = "Março"
	case 4:
		mesExtenso = "Abril"
	case 5:
		mesExtenso = "Maio"
	case 6:
		mesExtenso = "Junho"
	case 7:
		mesExtenso = "Julho"
	case 8:
		mesExtenso = "Agosto"
	case 9:
		mesExtenso = "Setembro"
	case 10:
		mesExtenso = "Outubro"
	case 11:
		mesExtenso = "Novembro"
	case 12:
		mesExtenso = "Dezembro"
	default:
		fmt.Println("Informe um mês válido!")
	}
	fmt.Printf("Data: %d de %s de %d\n", dia, mesExtenso, ano)

}
