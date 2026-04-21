package main

import "fmt"

func main() {
	var n int

	soma := 0
	quantidade := 0
	maior := 0
	menor := 0

	somaPares := 0
	contPares := 0
	contImpares := 0

	for {
		fmt.Print("Digite um número (30000 para sair): ")
		fmt.Scan(&n)

		if n == 30000 {
			break
		}

		if quantidade == 0 {
			maior = n
			menor = n
		}

		soma += n
		quantidade++

		if n > maior {
			maior = n
		}

		if n < menor {
			menor = n
		}

		if n%2 == 0 {
			somaPares += n
			contPares++
		} else {
			contImpares++
		}
	}

	fmt.Println("Soma:", soma)
	fmt.Println("Quantidade:", quantidade)

	if quantidade > 0 {
		fmt.Println("Média:", float64(soma)/float64(quantidade))
	}

	fmt.Println("Maior:", maior)
	fmt.Println("Menor:", menor)

	if contPares > 0 {
		fmt.Println("Média dos pares:", float64(somaPares)/float64(contPares))
	}

	if quantidade > 0 {
		fmt.Println("Percentual de ímpares:",
			(float64(contImpares)/float64(quantidade))*100, "%")
	}
}
