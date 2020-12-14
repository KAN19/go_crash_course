package main

import "fmt"

func main() {
	ids := []int{32, 34, 12, 25, 65}

	// Loop throug ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	// Add Ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d", sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Mike": "Mike@gmail.com", "Sharon": "Sharon@gmail.com"}

	// k : key, v : value
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name" + k)
	}
}
