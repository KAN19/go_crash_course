package main

import "fmt"

func main() {
	// 	// Ddefine map
	// 	emails := make(map[string]string)

	// 	// Assign
	// 	emails["Bob"] = "bobtabteasd@gmail.com"
	// 	emails["Tom"] = "TomHaybea@gmail.com"
	// 	emails["Mike"] = "Mkiene@gmail.com"

	//Declare map and add kv
	emails := map[string]string{"Bob": "Bob@gmail.com", "Sharon": "Sharon@gmail.com"}

	emails["Mike"] = "Mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
