package main

import (
	"fmt"
	"strings"
)

type Posisi struct {
	gaji      int
	scoreTest int
	plusGaji  int
}

func main() {
	var applicant string
	applicant = "Fajrin"
	var code string
	code = "aBdfmjn"
	var experience int
	experience = 5
	var scoreTest int
	scoreTest = 90
	var posisi string
	var salary int

	frontend := Posisi{
		gaji:      11000000,
		scoreTest: 90,
		plusGaji:  1000000,
	}
	backend := Posisi{
		gaji:      10000000,
		scoreTest: 85,
		plusGaji:  1000000,
	}
	if c := strings.Index(code, "B"); c > -1 {
		posisi = "Backend Developer"
		status := ""
		if experience > 4 {
			status = "Sr"
			salary += 1000000
		} else {
			status = "Jr"
		}
		salary += backend.gaji
		if scoreTest > backend.scoreTest {
			salary += backend.plusGaji
		}
		fmt.Printf("Selamat %s kamu diterima menjadi %s. %s dengan gaji pokok sebesar %d", applicant, status, posisi, salary)
	} else {
		posisi = "posisi tidak ditemukan"
	}

	if c := strings.Index(code, "F"); c > -1 {
		posisi = "Frontend Developer"
		status := ""
		if experience > 4 {
			status = "Sr"
			salary += 1000000
		} else {
			status = "Jr"
		}
		salary += frontend.gaji
		if scoreTest > frontend.scoreTest {
			salary += frontend.plusGaji
		}
		fmt.Printf("Selamat %s kamu diterima menjadi %s. %s dengan gaji pokok sebesar %d", applicant, status, posisi, salary)
	} else {
		posisi = "posisi tidak ditemukan"
	}

}
