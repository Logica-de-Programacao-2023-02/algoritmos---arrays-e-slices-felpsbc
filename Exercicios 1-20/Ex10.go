package main

import "fmt"

func main() {
	slice := []int{10, 12, 16, 24, 11, 165, 241, 143, 43, 23}
	var max, min int

	if len(slice) == 0 {
		fmt.Println("O slice está vazio.")
	}

	max = slice[0]
	min = slice[0]

	for _, i := range slice {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}

	fmt.Println("O valor máximo é:", max)
	fmt.Println("O valor mínimo é:", min)
}
