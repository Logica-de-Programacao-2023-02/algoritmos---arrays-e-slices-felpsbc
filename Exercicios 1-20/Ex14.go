package main

import "fmt"

func main() {

	Slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var indice1, indice2 int
	fmt.Println("Escreva os índices dos números que deverão ser trocados de posição ")
	fmt.Scan(&indice1, &indice2)

	if indice1 < 0 || indice1 > 7 || indice2 < 0 || indice2 > 7 {
		fmt.Println("Indices invalidos")
		return
	}
}
