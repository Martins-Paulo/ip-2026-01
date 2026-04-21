package main

import "fmt"

func main() {
	var idade int
	var altura, peso float64
	var opcao int

	totalPessoas := 0
	mais50 := 0
	somaAltura := 0.0
	cont1020 := 0
	menos40kg := 0

	for {
		fmt.Print("Idade: ")
		fmt.Scan(&idade)

		fmt.Print("Altura: ")
		fmt.Scan(&altura)

		fmt.Print("Peso: ")
		fmt.Scan(&peso)

		totalPessoas++

		if idade > 50 {
			mais50++
		}

		if idade >= 10 && idade <= 20 {
			somaAltura += altura
			cont1020++
		}

		if peso < 40 {
			menos40kg++
		}

		fmt.Print("Deseja continuar? (1 = sim): ")
		fmt.Scan(&opcao)

		if opcao != 1 {
			break
		}
	}

	fmt.Println("Pessoas com mais de 50 anos:", mais50)

	if cont1020 > 0 {
		fmt.Println("Média das alturas (10 a 20 anos):", somaAltura/float64(cont1020))
	}

	fmt.Println("Porcentagem abaixo de 40kg:",
		(float64(menos40kg)/float64(totalPessoas))*100, "%")
}
