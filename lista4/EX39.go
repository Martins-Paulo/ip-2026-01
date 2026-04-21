package main

import "fmt"

func main() {
	var id, peso int

	var maiorPeso, menorPeso int
	var idMaior, idMenor int

	for i := 0; i < 90; i++ {
		fmt.Scan(&id, &peso)

		if i == 0 || peso > maiorPeso {
			maiorPeso = peso
			idMaior = id
		}

		if i == 0 || peso < menorPeso {
			menorPeso = peso
			idMenor = id
		}
	}

	fmt.Println("Mais gordo:", idMaior, maiorPeso)
	fmt.Println("Mais magro:", idMenor, menorPeso)
}
