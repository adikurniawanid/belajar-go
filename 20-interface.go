package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hashName HasName) {
	fmt.Println("Hello", hashName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var adi Person
	adi.Name = "Adi"

	sayHello(adi)

	var cat Animal
	cat.Name = "Zero Pus"

	sayHello(cat)
}
