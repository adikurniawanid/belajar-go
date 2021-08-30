package main

import "fmt"

// mirip kayak java yang semester 3 itu loh, lupa namanya hehe
type Mahasiswa struct {
	Nama, NIM string
	Angkatan  int
}

func main() {
	var adi Mahasiswa
	adi.Nama = "Adi Kurniawan"
	adi.NIM = "0902"
	adi.Angkatan = 2018

	fmt.Println(adi)
	fmt.Println(adi.Nama)
	fmt.Println(adi.Angkatan)
	fmt.Println(adi.NIM)

	ida := Mahasiswa{
		Nama:     "Ida Kurnia",
		NIM:      "0901",
		Angkatan: 2018,
	}

	fmt.Println(ida)
	fmt.Println(ida.Nama)
	fmt.Println(ida.Angkatan)
	fmt.Println(ida.NIM)

	edi := Mahasiswa{"Edi", "0903", 2017} //versi ini harus urutan
	fmt.Println(edi)
	fmt.Println(edi.Nama)
	fmt.Println(edi.Angkatan)
	fmt.Println(edi.NIM)
}
