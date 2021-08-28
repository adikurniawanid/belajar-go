package main

import "fmt"

func main() {
	// tipedata slice memiliki 3 data,  pointer, length, capacity
	// pointer adalah pentujuk data pertama diarray para slice
	// length adalah panjang dari slice
	// capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	// array[1:10] membuat slice dari index array 1 sampe 9
	// array[1:] membuat slice dari index array 1 sampai seterusnya
	// array[:10] membuat slice dari index array 0 hingga sebelum 10 atau 9
	// array[:] membuat slice dari seluruh array

	// catatan mengubah array mempengaruhi slice

	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(bulan[4])
	var slice1 = bulan[4:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	bulan[5] = "diubah"
	fmt.Println(slice1)
	fmt.Println("kapasitas", cap(slice1))

	var hari = []string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	// catatan mengubah slice mempengaruhi array asli

	hariSlice := hari[5:]
	hariSlice[0] = "Sabtu Kece"
	hariSlice[1] = "Minggu Kece"
	fmt.Println(hari)

	// ketika capacity masih cukup akan menggunakan array lama
	// tapi kalau capacity tidak cukup akan menggunakan baru

	hariSlice2 := append(hariSlice, "Libur Baru")
	hariSlice2[0] = "ups"
	fmt.Println(hariSlice2)
	fmt.Println(hari)

	// slice dari awal tanpa array
	newSlice := make([]string, 2, 5) //length 2, capacity 5
	newSlice[0] = "Adi"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	sliceC := make([]string, len(newSlice), cap(newSlice))
	copy(sliceC, newSlice) //ke slice mana, dari slice mana

	fmt.Println(sliceC)

	// array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
