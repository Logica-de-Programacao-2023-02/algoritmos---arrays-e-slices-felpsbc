package main

import "fmt"

func main() {

	array := [4]float64{1.5, 2.7, 6.4, 1.1}
	fmt.Println("Escreva 4 números")
	fmt.Scan(&array)

	var multiplicação float64
	multiplicação = 1

	for _, num := range array {
		multiplicação *= num
	}
	fmt.Printf("A multiplicação dos números é: %.2f", multiplicação)
}
