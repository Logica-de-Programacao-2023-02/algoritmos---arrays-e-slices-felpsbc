package main

import "fmt"

func main() {

	var lista = []int{1, 2, 3, 4, 5}
	lista = append(lista[:2], lista[3:]...)
	fmt.Println(lista)
}
