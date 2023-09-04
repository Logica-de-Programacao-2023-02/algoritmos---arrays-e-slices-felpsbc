package main

import "fmt"

func main() {

	matriz := [3][2]int{}

	for i := 0; i < 3; i++ {
		for c := 0; c < 2; c++ {
			fmt.Println("Informe os valores das posições", i, c)
			fmt.Scan(&matriz[i][c])
		}
	}
	fmt.Println("Resultado da matriz:")
	for i := 0; i < 3; i++ {
		for c := 0; c < 2; c++ {
			fmt.Printf("%d", matriz[i][c])
		}
		fmt.Println()
	}

}
