package main

import "fmt"

func main() {
	var fahrenheit, celsius, qtdChuva, qtdMilimetro float64

	fmt.Print("Informe a temperatura em Fahrenheit e a quantidade de chuva em polegadas: ")
	fmt.Scan(&fahrenheit, &qtdChuva)

	celsius = ((5.0 * fahrenheit) - 160) / 9
	qtdMilimetro = qtdChuva * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f\n", celsius, qtdMilimetro)

}
