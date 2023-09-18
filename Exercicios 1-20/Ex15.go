package main

import "fmt"

func main() {

	array := [10]float64{1.0, 1.5, 2.0, 2.5, 3.0, 5.0, 6.0, 6.5, 7.8, 9.9}
	slice := []float64{}

	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}
	fmt.Println("Novo Slice: ", slice)
}
