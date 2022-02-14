package main

import (
	"fmt"
	"strings"
)

type BanyakAngka struct {
	angka       string
	jumlahAngka int
}

func main() {

	word := "Aku adalah Ak"
	word = strings.ToLower(word)
	var maxAngka BanyakAngka
	for _, v := range word {
		totalAngka := 0
		for _, y := range word {
			if v == y {
				totalAngka += 1
			}
		}
		if totalAngka > maxAngka.jumlahAngka {
			maxAngka.jumlahAngka = totalAngka
			maxAngka.angka = string(v)
		}
	}
	if len(word) == 1 {
		fmt.Printf("Hanya memiliki satu karakter yaitu %s", word)
	} else if maxAngka.jumlahAngka > 10 {
		fmt.Printf("Karakter terbanyak adalah %s jumlah yang sangat banyak yaitu %d", maxAngka.angka, maxAngka.jumlahAngka)
	} else if maxAngka.jumlahAngka > 5 && maxAngka.jumlahAngka < 10 {
		fmt.Printf("Karakter terbanyak adalah %s dengan jumlah yang cukup banyak yaitu %d", maxAngka.angka, maxAngka.jumlahAngka)
	} else if maxAngka.jumlahAngka <= 5 {
		fmt.Printf("Karakter terbanyak adalah %s dengan jumlah %d", maxAngka.angka, maxAngka.jumlahAngka)
	}

}
