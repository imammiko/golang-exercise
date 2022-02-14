package main

import "fmt"

func main() {
	var age int
	// age = 3
	age = 8
	var height int
	// height = 95
	height = 140
	var output int
	switch {
	case age < 1:
		fmt.Println("dilarang masuk")
	case age >= 2 && age <= 3:
		output += 30000
		if height > 70 {
			output += 10000
		}
	case age >= 4 && age <= 7:
		output += 40000
		if height > 120 {
			output += 120000
		}
	case age >= 8 && age <= 10:
		output += 50000
		if height > 150 {
			output += 20000
		}
	case age > 10:
		output += 80000
	}
	fmt.Println(output)
}
