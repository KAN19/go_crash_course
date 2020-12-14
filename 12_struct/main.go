package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method with value reciever
func (p Person) greet() string {
	return "Hello, my name is " + p.lastName + " " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

// getMaried pointer reciever
func (p *Person) getMaried(spouseLastName string) {
	if p.gender == "Nam" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Kiet", lastName: "Nguyen", city: "Ho Chi Minh", gender: "Nam", age: 20}
	person2 := Person{firstName: "Samathar", lastName: "Nguyen", city: "Ho Chi Minh", gender: "Nu", age: 20}

	person1.hasBirthday()
	person2.getMaried("William")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
