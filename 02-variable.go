package main

import "fmt"

func main() {
	var namaDepan string
	// alternatif tanpa perlu menyebutkan var
	umur := 21
	namaDepan = "Adi"
	fmt.Println(namaDepan)

	// kalau diisi langsung, tidak perlu menyebut tipe data
	var namaBelakang = "Kurniawan"
	fmt.Println(namaBelakang)

	fmt.Println(umur)

	// alternatif kalau mau sekaligus
	var (
		firstName = "adi"
		lastName  = "kurniawan"
	)

	fmt.Println(firstName, lastName)
}
