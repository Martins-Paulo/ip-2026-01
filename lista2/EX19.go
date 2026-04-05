// Desenvolver um programa com as opções de calcular e imprimir o volume e a área da superfície de um cone
// reto, de um cilindro ou de uma esfera. O programa deverá ler a opção da figura desejada (1-cone / 2-cilindro /
// 3-esfera) e de acordo com a opção escolhida calcular e escrever o volume e a área da superfície da figura
// pedida.

package main

import (
	"fmt"
	"math"
)

func main() {

	var raio, altura, volume, area float64
	var forma int

	fmt.Println("Escolha a forma geometrica para calcular área: \n1 - Cone Reto\n2 - Cilindro\n3 - Esfera.")
	fmt.Scan(&forma)

	fmt.Println("Informe, caso possua, o valor do raio e altura da figura:")
	fmt.Scan(&raio, &altura)

	switch forma {
	case 1:
		volume = (math.Pi * math.Pow(raio, 2) * altura) / 3
		area = (math.Pi * raio * math.Sqrt((math.Pow(raio, 2) + math.Pow(altura, 2))))
	case 2:
		volume = (math.Pi * math.Pow(raio, 2) * altura)
		area = 2 * math.Pi * raio * altura

	case 3:
		volume = (4 / 3.0) * math.Pi * math.Pow(raio, 3)
		area = 4 * math.Pi * math.Pow(raio, 2)
	}

	fmt.Printf("Forma número %d\nVolume: %.2f m³\nÁrea: %.2f m²\n", forma, volume, area)

}
