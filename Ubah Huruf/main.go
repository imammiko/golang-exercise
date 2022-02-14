package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata = "i love javascript"
	kataSlice := strings.Split(kata, "")
	var vowel = "aiueo"
	for i, v := range kataSlice {
		if strings.Index(vowel, v) > -1 {
			kataSlice[i] = "$"
		}
	}
	kata = strings.Join(kataSlice, "")
	fmt.Println(kata)
}
