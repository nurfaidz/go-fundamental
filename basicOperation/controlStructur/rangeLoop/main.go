package main

import "fmt"

func main() {
	cityCodes := []string{"CGK", "BDO", "YOG", "XYZ"}

	for _, kodeKota := range cityCodes {
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

		fmt.Println("Kode kota: %s, Nama kota: %s\n", kodeKota, namaKota)
	}
}
