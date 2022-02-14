package main

import (
	"fmt"
	"strconv"
)

func main() {
	var route string
	route = "123456"
	var output string
	for _, v := range route {
		output += "|"
		for i := 1; i <= 6; i++ {
			fmt.Println(string(v), i)
			number, _ := strconv.Atoi(string(v))
			if number == i {
				output += "o"
			} else {
				output += "#"
			}
		}
		output += "|"
		output += "\n"
	}
	fmt.Println(output)

}
