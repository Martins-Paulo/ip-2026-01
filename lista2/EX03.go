package main

//3. Elabore um programa que leia dois números inteiros e efetue a adição dos mesmos. Caso o valor
//somado seja maior do que 20, o resultado a ser apresentado será a soma mencionada adicionada de
//8. Caso o valor somado seja menor ou igual a 20, deve-se subtrair 5 do mesmo para apresentá-lo em
//seguida.

import "fmt"

func main() {
	numero1, numero2, soma := 0, 0, 0

	fmt.Print("Informe primeiro e o segundo números: ")
	fmt.Scan(&numero1, &numero2)

	soma = numero1 + numero2

	if soma > 20 {
		soma = soma + 8
		fmt.Printf("Soma é %d!\n", soma)
	} else {
		soma = soma - 5
		fmt.Printf("Soma é %d!\n", soma)
	}

}
