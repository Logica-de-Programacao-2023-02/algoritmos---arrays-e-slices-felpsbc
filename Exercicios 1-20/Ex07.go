package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	var num int
	fmt.Println("Escreva um digito: ")
	fmt.Scan(&num)

	x := false

	for _, numero := range slice {
		if num == numero {
			x = true
			break
		}
	}
	if x == false {
		slice = append(slice, num)
	}
	fmt.Println(slice)
}
