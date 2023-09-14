package main

import "fmt"

func main() {

	array := [4]float64{}
	for i := 0; i < len(array); i++ {
		fmt.Printf("\n Escreva o número %d da lista: ", i+1)
		fmt.Scan(&array[i])
	}

	var multiplicação float64
	multiplicação = 1

	for _, num := range array {
		multiplicação *= num
	}
	fmt.Printf("A multiplicação dos números é: %.2f", multiplicação)
}
