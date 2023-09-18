package main

import "fmt"

func main() {

	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num int
	fmt.Println("Escreva um número que será adicionado ao primeiro e último números do array:")
	fmt.Scan(&num)

	array[0] += num
	array[6] += num

	fmt.Println("O novo array é:", array)

}
