package main

import "fmt"

func main() {
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Sua Matriz:")
	for l := 0; l < 2; l++ {
		for c := 0; c < 3; c++ {
			fmt.Printf("%d", matriz[l][c])
		}
		fmt.Println()
	}

	var linha, coluna int

	fmt.Printf("Informe o índice de linha: %d\n", linha)
	fmt.Scan(&linha)
	fmt.Printf("Informe o índice de coluna: %d\n", coluna)
	fmt.Scan(&coluna)

	fmt.Println("O seu numero é:", matriz[linha][coluna])
}
