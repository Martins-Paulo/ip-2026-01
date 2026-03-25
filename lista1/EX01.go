package main

import f "fmt"

func main() {

	nota1, nota2, nota3, media := 0.0, 0.0, 0.0, 0.0

	f.Print("Informe as nota 1, nota 2 e nota 3 do aluno(a): ")
	f.Scan(&nota1, &nota2, &nota3)

	media = (nota1 + nota2 + nota3) / 3

	if media >= 6 {
		f.Printf("Média: %.2f\nAluno Aprovado! :D.", media)
	} else {
		f.Printf("Média: %.2f\nAluno Reprovado! :C.", media)
	}

}
