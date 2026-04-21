package main

import "fmt"

func main() {
	var n int

	fmt.Print("Quantidade de alunos: ")
	fmt.Scan(&n)

	aprovados := 0
	exame := 0
	reprovados := 0

	somaClasse := 0.0

	for i := 0; i < n; i++ {
		var nota1, nota2 float64

		fmt.Print("Nota 1: ")
		fmt.Scan(&nota1)

		fmt.Print("Nota 2: ")
		fmt.Scan(&nota2)

		media := (nota1 + nota2) / 2
		somaClasse += media

		fmt.Println("Média:", media)

		if media <= 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	fmt.Println("Aprovados:", aprovados)
	fmt.Println("Exame:", exame)
	fmt.Println("Reprovados:", reprovados)

	fmt.Println("Média da classe:", somaClasse/float64(n))
}
