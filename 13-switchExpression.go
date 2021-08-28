package main

import "fmt"

func main() {
	name := "adi"

	switch name {
	case "adsi":
		fmt.Println("Hello adi")
	case "kurniawan":
		fmt.Println("Halo kurniawan")
	default:
		fmt.Println("Halo bukan siapa siapa")
	}

	// short statement di switch

	namaBelakang := "Kurniawan"

	switch length := len(namaBelakang); length > 5 {
	case true:
		fmt.Println("contoh aja")
	case false:
		fmt.Println("gitu deh")
	}

	// tanpa kondisi dengan kondisi didalam case

	length := len(namaBelakang)
	switch {
	case length > 10:
		fmt.Println("nama panjang")
	case length < 10:
		fmt.Println("nama pendek")
	default:
		fmt.Println("nama pas")
	}
}
