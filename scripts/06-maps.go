// This is the namespace to tell Go compiles that this is a standalone executable program. The main package is the entry point of the program.
package main

import (
	"fmt"
	"maps"
)

// The main function is the entry point of the program. When you run the program, the code inside the main function will be executed.
func main() {
    m := make(map[string]int)

		m["apple"] = 5
		m["banana"] = 3

		fmt.Println("Map:", m)
		fmt.Println("Length map", len(m))
		fmt.Println("Value for 'apple':", m["apple"])
		fmt.Println("Value for 'banana':", m["banana"])

	// declare and initialize a map in a single line.
		 n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

		  n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}