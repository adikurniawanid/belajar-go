package main

import "fmt"

func main() {
	//ini array arr := [...]int
	//ini slice arr := []int
	var nama [2]string
	nama[0] = "Adi"
	nama[1] = "Kurniawan"

	fmt.Println(nama[0])
	fmt.Println(nama[1])

	var values = [3]int{
		90, 80,
	}

	// panjang array bukan banyak data
	fmt.Println(len(values), len(nama))

	fmt.Println(values)

	var nilai = [...]int{
		2, 3, 4, 5,
	}
	fmt.Println(len(nilai))

	fmt.Println(nilai)

}
