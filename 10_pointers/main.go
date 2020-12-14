package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T", b)

	fmt.Print(*b)
	fmt.Print(*&a)

	*b = 10
	fmt.Println(a)
}
