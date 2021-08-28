package main

import "fmt"

func main() {
	var a = 20
	var b = 10
	var c = a + b
	fmt.Println(c)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a += 20
	fmt.Println(a)
	a++
	fmt.Println(a)

	var status = true
	fmt.Println(!status)
}
