package main

import "fmt"

func main() {

	var tamanho int
	fmt.Println("Escreva o tamanho do tamanho da lista")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("Digite os números")
		fmt.Scan(&slice[i])
	}

	var soma int

	for _, num := range slice {
		soma += num
	}
	fmt.Println("A soma dos valores adicionados é:", soma)

}
