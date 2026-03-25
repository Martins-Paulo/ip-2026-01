package main

import (
	"fmt"
	"math"
)

func main() {

	const pi float64 = 3.14159
	var areaTotal, areaCirculo, areaLateral, custo, raio, altura float64

	fmt.Print("Informe o valores do raio e altura da lata: ")
	fmt.Scan(&raio, &altura)

	areaCirculo = pi * math.Pow(raio, 2)
	areaLateral = 2 * pi * raio * altura
	areaTotal = (2 * areaCirculo) + areaLateral

	custo = areaTotal * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)

}
