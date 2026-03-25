package main

import "fmt"

func main() {
	var quantidadesTemperaturas int64
	var fahrenheit, celsius float64

	fmt.Print("Informe a quantidades de temperaturas que serão convertidas: ")
	fmt.Scan(&quantidadesTemperaturas)

	for i := 0; i < int(quantidadesTemperaturas); i++ {

		fmt.Printf("Informe a temperatura %d: ", i+1)
		fmt.Scan(&fahrenheit)

		celsius = (5.0 * (fahrenheit - 32)) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)

	}

}
