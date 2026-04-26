package main

import (
	"fmt"
	"math"
)

// Constants with Capital letter are exported and can be accessed from other packages.
const Version string = "1.0.0"
// lowercase letter constants are unexported and can only be accessed within the same package. They are not visible to other packages.
const version2 = "1.0.0" // This constant is unexported because it starts with a lowercase letter.
func main() {
	fmt.Println("Version:", Version)
	// Function scope variables with lowercase letter are unexported and can only be accessed within the same package.
	const n = 42
	fmt.Println("n:", n)
	fmt.Println("Version2:", version2)

	// allowing defining a variable without specifying the type, Go will infer the type based on the value assigned to it.
 	var f string = "apple"
	fmt.Println("Version2:", version2)
	fmt.Println("Fruit:", f)

	const Pi = math.Pi
	fmt.Println("Pi:", Pi)

	const n100 = 100 
 fmt.Println("n100:", n100)

 // All math ops expects a float64
 nSum := math.Pow(n, n100)
 fmt.Println("nSum:", nSum)

 var nInt int = 2
 var nFloat = float64(nInt)
 nSumInt := math.Pow(n, nFloat)
 fmt.Println("nSumInt:", nSumInt)

}
