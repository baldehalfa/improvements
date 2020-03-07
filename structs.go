package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Mariame", lastName: "Diallo", city: "Conakry", gender: "f", age: 21}
	// Alternative
	person2 := Person{"Alpha", "Balde", "Kuala Lumpur", "m", 30}

	person1.hasBirthday()
	person1.getMarried("Diallo")

	person2.getMarried("Barry")

	fmt.Println(person2.greet())
}
