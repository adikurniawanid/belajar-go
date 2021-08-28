package main

import "fmt"

func main() {
	// key nya string, valuenya string
	kuliah := map[string]string{
		"universitas": "UNSRI",
		"fakultas":    "Ilmu Komputer",
	}

	kuliah["jurusan"] = "Teknik Informatika"

	fmt.Println(kuliah)
	fmt.Println(len(kuliah))
	fmt.Println(kuliah["universitas"])

	// buat map tanpa data terlebuh dahulu
	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "dika"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
