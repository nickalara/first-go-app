package main

import (
	"fmt"
	"path"
)

// Leftover from first lesson where we called the name of the mascot from the mascot.go package
// func main() {
// 	fmt.Println(mascot.BestMascot())
// }

// This function demonstrates the Split function, how to declare 2 variables at the same time, and how to use 2 variables at the same time in a function call.
// func main() {
// 	var dir, file string
// 	dir, file = path.Split("css/main.css")

// 	fmt.Println("dir :", dir)
// 	fmt.Println("file name:", file)
// }

// This demonstrates how to discard values. In this case, notice that the first returned value is not used. We use a blank identifier (underscore) to achieve this.
// func main() {
// 	var file string
// 	_, file = path.Split("css/main.css")

// 	fmt.Println("file name:", file)
// }

// This demonstrates short declarations. Note that we can declare and assign the value of file in the same line of code using the ":=" operator.
func main() {
	_, file := path.Split("css/main.css")

	fmt.Println("file name:", file)
}
