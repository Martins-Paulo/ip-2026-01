package main

import (
	"fmt"
	"math"
)

func main() {

	var n1, n2, n3, centena, dezena, unidade, numeroCalculado, numeroQuadrado int32

	fmt.Print("Informe os 3 números: ")
	fmt.Scan(&n1, &n2, &n3)

	if n1 < 10 && n2 < 10 && n3 < 10 {
		centena = n1 * 100
		dezena = n2 * 10
		unidade = n3 * 1

		numeroCalculado = centena + dezena + unidade
		numeroQuadrado = int32(math.Pow(float64(numeroCalculado), 2))

		fmt.Print(numeroCalculado, numeroQuadrado)

	} else {
		fmt.Print("DIGITO INVALIDO")
	}

}
