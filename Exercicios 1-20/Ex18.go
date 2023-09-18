package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Digite a quantidade de números primos que deseja visualizar: ")
	fmt.Scan(&n)

	lista := make([]int, 0, n)
	for i := 2; len(lista) < n; i++ {
		primo := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo == true {
			lista = append(lista, i)
		}
	}
	fmt.Printf("Os %d primeiros números primos são: %v\n", n, lista)

}
