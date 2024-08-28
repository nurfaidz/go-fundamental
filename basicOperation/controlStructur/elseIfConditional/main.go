package main

import "fmt"

func main() {
	kodeKota := "YOG"

	var namaKota string
	if kodeKota == "CGK" {
		namaKota = "Jakarta"
	} else if kodeKota == "BDO" {
		namaKota = "Bandung"
	} else if kodeKota == "YOG" {
		namaKota = "Yogyakarta"
	} else {
		namaKota = "Kode kota tidak valid!"
	}

	fmt.Println(namaKota)
}
