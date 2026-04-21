package main

import "fmt"

func main() {
	var cpf string
	fmt.Scan(&cpf)

	var nums [11]int

	for i := 0; i < 11; i++ {
		nums[i] = int(cpf[i] - '0')
	}

	// primeiro dígito
	soma := 0
	peso := 10
	for i := 0; i < 9; i++ {
		soma += nums[i] * peso
		peso--
	}

	resto := soma % 11
	d1 := 0
	if resto >= 2 {
		d1 = 11 - resto
	}

	// segundo dígito
	soma = 0
	peso = 11
	for i := 0; i < 10; i++ {
		soma += nums[i] * peso
		peso--
	}

	resto = soma % 11
	d2 := 0
	if resto >= 2 {
		d2 = 11 - resto
	}

	if d1 == nums[9] && d2 == nums[10] {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}