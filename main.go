package main

import "fmt"

func main() {
	helloWorld := "hello, world!"
	fmt.Println(helloWorld)

	// Lesson 6: Type Conversion
	speed := 100 // speed is int (whole numbers only)
	force := 2.5 // force is float64 (decimals or partial numbers)

	// Cannot multiply mismatched types (int and float64 in this case)
	// speed = speed * force // Compiler error
	power := speed * int(force)    // We can convert the type of force to match the type of speed
	fmt.Println(power)             // But this result is not ideal... If you think about it, the power should actually be 250...
	fmt.Println(force, int(force)) // It's important to note that changing the type can therefore also change the value, but not the variable itself... So be careful...

	// // Trying to convert speed to int will return an error. This is because we previously declared power as an int.
	// power = float64(speed) * force
	// // Rather than going back, let's make a new variable to demonstrate the best approach.
	power2 := int(float64(speed) * force) // Think about the order of the calculation. Similar in math to the order of operations, the sequence of events follows logically. Keep this in mind.
	fmt.Println(power2)
}
