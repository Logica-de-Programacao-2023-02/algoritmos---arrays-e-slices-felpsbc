package main

import "fmt"

func main() {

	array := [6]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	var num float64
	fmt.Printf("Digite um número que será adicionado a todos os outros dentro do array\n")
	fmt.Scan(&num)

	for i := range array {
		array[i] += num
	}
	fmt.Println(array)
}
