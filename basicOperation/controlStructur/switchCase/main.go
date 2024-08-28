package main

import "fmt"

func main() {
	kodeKota := "CGK"

	var namaKota string
	switch kodeKota {
	case "CGK":
		namaKota = "Jakarta"
	case "BDO":
		namaKota = " Bandung"
	case "YOG":
		namaKota = "Yogyakarta"
	default:
		namaKota = "Kode kota tidak valid!"
	}

	fmt.Println(namaKota)
}
