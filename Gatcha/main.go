// Generating Random Numbers in Golang
package main

import (
	"math/rand"
	"time"
)

func main() {

	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	value := y1.Intn(4) + 1
	switch value {
	case 1:
		println("coba lagi ya")
	case 2:
		println("selamat kamu mendapatkan kupon sebanyak 5")
	case 3:
		println("selamat kamu mendapatkan kupon sebanyak 15")
	case 4:
		println("selamat kamu mendapatkan kupon sebanyak 50")
	case 5:
		println("WOW, kamu menang jackpot! Selamat!!")
	}

}
