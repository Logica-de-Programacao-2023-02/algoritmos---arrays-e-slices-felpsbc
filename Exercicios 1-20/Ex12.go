package main

import "fmt"

func main() {

	var array = [5]int{3, 12, 10, 21, 35}

	slice := []int{}
	for _, i := range array {
		if i%3 == 0 {
			slice = append(slice, i)
		}

	}
	fmt.Println(slice)

}
