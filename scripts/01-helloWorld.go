// This is the namespace to tell Go compiles that this is a standalone executable program. The main package is the entry point of the program.
package main

// fmt is the based package for I/O in Go. It provides functions for formatting and printing output to the console.
import "fmt"

// The main function is the entry point of the program. When you run the program, the code inside the main function will be executed.
func main() {
    fmt.Println("hello world")
}