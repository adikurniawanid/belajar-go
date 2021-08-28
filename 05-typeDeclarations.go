package main

import "fmt"

func main() {
	// mengaliaskan tipe data yang sudah ada
	type NIM string
	type married bool

	var adiNIM NIM = "09021181823168"
	var adiMarriedStatus married = false
	fmt.Println(adiNIM, adiMarriedStatus)
}
