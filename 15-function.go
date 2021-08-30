package main

import (
	"fmt"
)

func sayHelloWorld(namaDepan string, namaBelakang string) {
	fmt.Println("Halo Dunia", namaDepan, namaBelakang)
}

func exForFuncWithReturn() string {
	return "Adi Kurniawan"
}

func exForFuncWithMultiReturn() (string, string) {
	return "Adi", "Kurniawan"
}

func exForFuncWithNamedReturn() (firstName, middleName, lastName string) {
	firstName = "Adi"
	middleName = "Kurnia"
	lastName = "wan"
	return
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good bye," + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(nama string) string {
	if nama == "anjing" {
		return "..."
	} else {
		return nama
	}
}

// type alias
type Filter func(string) string

func sayHelloWithFilterLagi(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// anonymous function
type BlackList func(string) bool

func registerUser(name string, BlackList BlackList) {
	if BlackList(name) {
		fmt.Println("Kamu keblok")
	} else {
		fmt.Println("welcome", name)

	}
}

// rekursif
func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	sayHelloWorld("adi", " kurniawan")
	result := exForFuncWithReturn()
	namaAwal, namaAkhir := exForFuncWithMultiReturn()

	fmt.Println(result)
	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)

	// menghiraukan value lain
	namaA, _ := exForFuncWithMultiReturn()
	_, namaB := exForFuncWithMultiReturn()
	fmt.Println(namaB)

	fmt.Println(namaA)

	a, b, c := exForFuncWithNamedReturn()
	fmt.Println(a, b, c)

	slice := []int{
		10, 123, 1231, 1231, 23,
	}

	total := sumAll(10, 21, 213, 123, 4124, 123, 21)
	fmt.Println(total)
	total = sumAll(slice...)
	fmt.Println(total)

	// function value
	goodbye := getGoodBye
	fmt.Println(goodbye("Adi Kurniawan"))

	// function as parameter
	sayHelloWithFilter("Adi Kurniawan", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)

	// anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("adi", blacklist)
	registerUser("anjing", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	// rekursif
	fmt.Println(factorial(4))
}
