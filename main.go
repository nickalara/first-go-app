package main

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
// CONVENTION: DO NOT USE SHORT DECLARATION IF YOU DON'T KNOW THE INITIAL VALUE
// func main() {
// 	_, file := path.Split("css/main.css")

// 	fmt.Println("file name:", file)
// }

// func main() {
// 	// score := 0 // DONT DO THIS
// 	var score int  // When you initialize an int variable it is automatically set to 0. This is better practice.
// 	println(score) // CHeck here. console outputs score value as "0" 👍

// 	// When you want to declare multiple variables, it's nice to group them based on relatedness. See this example:
// 	// var (
// 	// 	// related:
// 	// 	video string

// 	// 	// closely related:
// 	// 	duration int
// 	// 	current  int
// 	// )
//}

// WHEN TO USE SHORT DECLARATION:
// In Go, short declaration is actually preferred wherever possible to keep code concise.
func main() {
	width, height := 100, 50 // Since we know both starting values and they are non-zero, use short. Note that you can assign multiple variables at once.

	// Use short declaration when you are redeclaring a variable. This can be combined with declaring a new variable as in the following:
	width, color := 50, "red"

	println(width, height, color)
}
