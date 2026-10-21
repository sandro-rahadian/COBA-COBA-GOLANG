package main

import "fmt"

type mobil struct {
	tipe      string
	cc        int
	kapasitas string
}

func luasJajarGenjang(alas, tinggi float64) float64 {
	return alas * tinggi
}

func main() {
	var mobil1 mobil

	mobil1.tipe = "Sport"
	mobil1.cc = 160
	mobil1.kapasitas = "4 Orang"

	fmt.Println(mobil1)

	// Luas jajar genjang = alas x tinggi
	var alas, tinggi float64
	fmt.Print("Masukkan alas: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	fmt.Printf("Luas jajar genjang: %.2f\n", luasJajarGenjang(alas, tinggi))
}
