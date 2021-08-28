package main

import "fmt"

func main() {
	// loops di golang cuma ada for

	// break untuk menghentikan perulangan
	// continue untuk skip perulangan ke perulangan selanjutnya

	for angka := 0; angka < 10; angka++ {
		fmt.Println(angka)
	}

	slice := []string{"Adi", "Kurniawan", "Gege", "Banget", "Hehe"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range untuk array, slice, map
	for i, value := range slice {
		if i == 2 {
			continue
		}
		fmt.Println("Index", i, "=", value)

		if i == 3 {
			break
		}
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	person := map[string]string{
		"nama awal":   "adi",
		"nama tengah": "kurnia",
		"nama akhir":  "wan",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
