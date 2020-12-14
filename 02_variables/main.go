package main

import "fmt"

func main() {
	// Using var
	// var name string = "Kiet"
	var age int = 20
	var isCool = true
	const isHot = false
	// var size float32 = 1.3

	//Shorthand
	name2 := "Kiettt"

	// name3, email := "Kiet", "Hello@gmail.com"

	fmt.Println(name2, age, isCool)
	fmt.Printf("%T", name2)
}
