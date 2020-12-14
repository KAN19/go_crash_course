package main

import "fmt"

func main() {
	// var fruiArr [2]string

	// fruiArr[0] = "Apple"
	// fruiArr[1] = "Orange"

	// fmt.Println(fruiArr)
	// fmt.Println(fruiArr[1])

	fruitArr := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)

	fmt.Println(len(fruitArr))

	fmt.Println(fruitArr[1:3])
}
