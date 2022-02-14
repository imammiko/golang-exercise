package main

import "fmt"

func main() {
	var nama string
	nama = "andika"
	var nilai int
	nilai = 100
	var grade string
	switch {
	case nilai >= 80 && nilai <= 100:
		grade = "A"
	case nilai >= 65 && nilai <= 79:
		grade = "B"
	case nilai >= 50 && nilai <= 64:
		grade = "C"
	case nilai >= 35 && nilai <= 49:
		grade = "D"
	case nilai >= 0 && nilai <= 34:
		grade = "E"
	}
	fmt.Printf("Nama: %s; score:%s", nama, grade)
}
