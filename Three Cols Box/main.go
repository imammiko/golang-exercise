package main

import "fmt"

func main() {
	fmt.Println(duaDimensi(5))
}

func duaDimensi(row int) [][]int {

	var number int
	slice := make([][]int, row)
	for i := 0; i < row; i++ {
		inner := make([]int, 3)
		for j := 0; j < 3; j++ {
			number += 1
			inner[j] = number
		}
		slice[i] = inner
	}

	return slice
}
