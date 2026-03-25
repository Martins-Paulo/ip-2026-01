package main

import "fmt"

func main() {
	var porcentagemKw, salario, qtdKwGasto, custoKw, custoConsumo, custoDesconto float64

	fmt.Print("Informe o valor do salario e a quantidade gasta de Kw: ")
	fmt.Scan(&salario, &qtdKwGasto)

	porcentagemKw = salario * 0.7
	custoKw = porcentagemKw / 100
	custoConsumo = (porcentagemKw * qtdKwGasto) / 100
	custoDesconto = custoConsumo - (custoConsumo * 0.1)

	fmt.Printf("Custo por Kw: %.2f\nCusto do consumo: %.2f\nCusto com desconto: %.2f", custoKw, custoConsumo, custoDesconto)

}
