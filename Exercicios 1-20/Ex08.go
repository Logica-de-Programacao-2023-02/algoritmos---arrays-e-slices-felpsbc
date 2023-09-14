package main

import "fmt"

func main() {

	var slice = []string{"pera", "uva", "banana", "pessego", "melao", "melancia", "mamao", "damasco"}

	var fruta string
	fmt.Println("Digite o nome de uma fruta:")
	fmt.Scan(&fruta)

	NovoSlice := []string{}

	for _, item := range slice {
		if item != fruta {
			NovoSlice = append(NovoSlice, item)
		}
	}
	slice = NovoSlice
	fmt.Println("Slice fica assim apos a retirada da fruta", fruta, ":", slice)
}
