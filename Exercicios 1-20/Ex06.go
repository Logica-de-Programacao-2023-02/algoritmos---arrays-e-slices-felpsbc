package main

import "fmt"

func main() {

	arr := [10]int{30, 103, 4, 17, 7, 8, 13, 12, 31, 19}

	var num int

	fmt.Println("Escreva um valor:")
	fmt.Scan(&num)

	x := false

	for _, elemento := range arr {
		if num == elemento {
			x = true
		}
	}

	if x == true {
		fmt.Println("O seu número foi encontrado no array")
	} else {
		fmt.Println("O seu número não foi encontrado no array")
	}
}
