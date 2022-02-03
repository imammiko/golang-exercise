package main

import (
	"fmt"
)

type GerbangKeluar struct {
	G1 int
	G2 int
	G3 int
}

func main() {
	var gerbangKeluar string
	gerbangKeluar = "IC Caruban"
	var golonganKendaraan string
	golonganKendaraan = "I"
	var saldo int
	saldo = 8000
	var tarif int
	IcCaruban := GerbangKeluar{
		G1: 8500,
		G2: 12500,
		G3: 12500,
	}
	IcNganjuk := GerbangKeluar{
		G1: 43500,
		G2: 66500,
		G3: 65500,
	}
	switch gerbangKeluar {
	case "IC Caruban":
		switch golonganKendaraan {
		case "I":
			tarif += IcCaruban.G1
		case "II":
			tarif += IcCaruban.G2
		case "III":
			tarif += IcCaruban.G3
		}
	case "IC Nganjuk":
		switch golonganKendaraan {
		case "I":
			tarif += IcNganjuk.G1
		case "II":
			tarif += IcNganjuk.G2
		case "III":
			tarif += IcNganjuk.G3
		}
	default:
		fmt.Println("Tarif tidak ditemukan")
	}
	if saldo < tarif {
		sisa := tarif - saldo
		fmt.Printf("Saldo anda kurang Rp:%d/ Tidak cukup untuk transaksi", sisa)
	} else {
		sisa := saldo - tarif
		fmt.Printf("saldo saldo anda adalah RP:%d semoga selamat sampai tujuan", sisa)
	}
}
