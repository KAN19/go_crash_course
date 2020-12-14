package main

import "fmt"

func main() {
	x := 13
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("Hello")
	}

	// swtich
	color := "asdasd"

	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("Default")
	}

}
