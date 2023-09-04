package main

import "fmt"

func main() {

	var numeros [3]int
	fmt.Println("Escreva 3 números ")
	fmt.Scan(&numeros[0], &numeros[1], &numeros[2])

	var soma int
	soma = numeros[0] + numeros[1] + numeros[2]

	fmt.Printf("A soma dos números do Array é: %d\n", soma)

}
