package main

import (
	"fmt"
	"math"

	strutil "github.com/NguyenKiet/go_crash_course/03_packages/strUltilies"
)

func main() {
	fmt.Println(math.Floor(20.8))
	fmt.Println(math.Ceil(3.4))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Hello ne"))

}
