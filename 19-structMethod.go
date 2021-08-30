package main

import "fmt"

type Mahasiswa struct {
	Nama, NIM string
	Angkatan  int
}

func (mahasiswa Mahasiswa) sayHello() {
	fmt.Println("Hello, My name is", mahasiswa.Nama)
}

func main() {
	adi := Mahasiswa{Nama: "Adi"}
	adi.sayHello()
}
