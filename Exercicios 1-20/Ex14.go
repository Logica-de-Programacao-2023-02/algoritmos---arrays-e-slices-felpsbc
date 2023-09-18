package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var indice1, indice2 int
	fmt.Println(slice)
	fmt.Println("Escreva os índices dos números que deverão ser trocados de posição ")
	fmt.Scan(&indice1, &indice2)

	if indice1 != indice2 {
		slice[indice1], slice[indice2] = slice[indice2], slice[indice1]
		fmt.Println("Novo Slice :", slice)
	} else {
		fmt.Println("Os indices são os mesmos então continua igual.", slice)
	}

}
