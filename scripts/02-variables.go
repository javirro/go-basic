package main

import "fmt"

// Variables with Capital letter are exported and can be accessed from other packages.
// In this case, Version is a global variable that can be accessed from other parts of the program or even from other packages.
var Version string = "1.0.0"
// lowercase letter variables are unexported and can only be accessed within the same package. They are not visible to other packages.
var version2 = "1.0.0" // This variable is unexported because it starts with a lowercase letter.
func main() {
	fmt.Println("Version:", Version)
	// Function scope variables with lowercase letter are unexported and can only be accessed within the same package.
	tempVersion := "2.0.0"

	// allowing defining a variable without specifying the type, Go will infer the type based on the value assigned to it.
 	var f string = "apple"
	fmt.Println("Temp Version:", tempVersion)
	fmt.Println("Version2:", version2)
	fmt.Println("Fruit:", f)

	Version = "1.0.1"
	fmt.Println("Updated Version:", Version)

	var j bool = true
	fmt.Println("Is Go fun?", j)
	var k int = 42
	fmt.Println("The answer to the Ultimate Question of Life, The Universe, and Everything:", k)

	var l float64 = 3.14
	fmt.Println("Pi:", l)

	var m complex128 = 1 + 2i
	fmt.Println("Complex number:", m)

}
