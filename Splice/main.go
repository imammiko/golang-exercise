package main

import "fmt"

func main() {
	fmt.Println(splice([]string{"idaz", "fajrin", "samir"}, "anggara", 1, 2))
}

func splice(data []string, add string, index int, del int) []string {
	return append(data[:index], append([]string{add}, data[index+del:]...)...)
}
