package main

import "fmt"

func main() {
	var nilai32 int32 = 3276821
	var nilai64 int64 = int64(nilai32)
	// terjadi overflow jika konversi ke nilai yang kecil dan tidak muat menampung
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var namaLengkap = "Adi Kurniawan"
	var hurufA byte = namaLengkap[0]
	var hurufAString string = string(hurufA)

	fmt.Println(namaLengkap)
	fmt.Println(hurufA)
	fmt.Println(hurufAString)
}
