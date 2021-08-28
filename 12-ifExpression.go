package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Adi" {
		fmt.Println("Hallo", name)
	} else if name == "Kurniawan" {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println("Hallo bukan siapa siapa")
	}

	// short statment
	if length := len(name); length > 3 {
		fmt.Println("panjang nama :", length, "; lebih dari 3")
	} else {
		fmt.Println("kurang cukup")
	}
}
