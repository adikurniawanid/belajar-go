package main

import "fmt"

func main() {
	// scope yang diatas bisa diakses oleh dibawah
	// scope yang dibawah tidak bisa diakses yang diatas
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)

	// salah satu cara
	name := "adi"

	inc := func() {
		name := "kurniawan"
		fmt.Println(name)
	}
	inc()
	fmt.Println(name)
}
