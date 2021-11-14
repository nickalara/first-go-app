package part1

import (
	"fmt"
	"path"

	"github.com/nickalara/first-go-app/mascot"
)

func part1() {
	// Lesson 1: Call another package
	// First lesson where we call the name of the mascot from the mascot.go package
	fmt.Println(mascot.BestMascot())

	// Lesson 1.5: Declaring Variables
	var item int
	item = 1
	fmt.Println(item)

	var item2 int
	item2 = 2
	fmt.Println(item2)

	var item3, item4 int
	item3 = 3
	item4 = 4
	fmt.Println(item3, item4)
	fmt.Println(item, item2, item3, item4)

	var item5, item6, item7 string
	item5 = "hello"
	item6 = ","
	item7 = "world!"
	fmt.Println(item5, item6, " ", item7)

	// Lesson 2: Split function - multiple variables and return values
	// This demonstrates the Split function, how to declare 2 variables at the same time, and how to use 2 variables at the same time in a function call.
	var dir, fileA string
	dir, fileA = path.Split("css/mainA.css")

	fmt.Println("dir :", dir)
	fmt.Println("file name:", fileA)

	// Lesson 3: Discard values
	// This demonstrates how to discard values. In this case, notice that the first returned value is not used. We use a blank identifier (underscore) to achieve this.
	var fileB string
	_, fileB = path.Split("css/mainB.css")

	fmt.Println("file name:", fileB)

	// Lesson 4: Short Declarations
	// This demonstrates short declarations. Note that we can declare and assign the value of file in the same line of code using the ":=" operator.
	// CONVENTION: DO NOT USE SHORT DECLARATION IF YOU DON'T KNOW THE INITIAL VALUE
	_, fileC := path.Split("css/mainC.css")

	fmt.Println("file name:", fileC)

	// Lesson 5: Short declarations Do's and Dont's
	// score := 0     // DONT DO THIS
	var score int      // When you initialize an int variable it is automatically set to 0. This is better practice.
	fmt.Println(score) // Check here. console outputs score value as "0" 👍

	// 	// When you want to declare multiple variables, it's nice to group them based on relatedness. See this example:
	var (
		// related:
		video string

		// closely related:
		duration int
		current  int
	)
	fmt.Println(video, duration, current)

	// WHEN TO USE SHORT DECLARATION:
	// In Go, short declaration is actually preferred wherever possible to keep code concise.
	width, height := 100, 50 // Since we know both starting values and they are non-zero, use short. Note that you can assign multiple variables at once.

	// Use short declaration when you are redeclaring a variable. This can be combined with declaring a new variable as in the following:
	width, color := 50, "red"

	fmt.Println(width, height, color)
}
