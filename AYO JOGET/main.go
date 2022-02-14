package main

import "fmt"

func main() {
	var exercise string
	exercise = "<>^v>>>"
	var userInput string
	userInput = "<>^^>><"
	var point int
	var persentase float32
	var kategori string
	// fmt.Println(string(userInput[0]), string(exercise[0]), point)
	if len(exercise) != len(userInput) {
		fmt.Println("panjang string tidak sama")
	}
	for i, v := range exercise {
		if string(v) == string(userInput[i]) {
			point += 1
		}
	}
	persentase = float32(point) / float32(len(exercise)) * 100
	switch {
	case persentase == 100:
		kategori = "Perfect"
	case persentase >= 80 && persentase <= 99:
		kategori = "great"
	case persentase >= 60 && persentase <= 79:
		kategori = "good"
	case persentase >= 0 && persentase <= 59:
		kategori = "Bad"
	}
	fmt.Println(point, persentase, kategori)
}
