package main

import "fmt"

func intSeq() func() int {
    i := 0
    // anonymous function that forms a closure. It captures the variable i from the surrounding scope and increments it each time it's called.
    return func() int {
        i++
        return i
    }
}

func main() {

    // intSeq returns a function that we assign to nextInt. This returned function is a closure that captures the variable i.
    nextInt := intSeq()

    fmt.Println(nextInt()) // Output: 1
    fmt.Println(nextInt()) // Output: 2
    fmt.Println(nextInt())  // Output: 3

    // Each call to intSeq creates a new instance of the variable i, so if we create another closure, it will have its own i.
    newInt2 := intSeq() // This new closure will have its own i, starting at 0.
    fmt.Println(newInt2()) // Output: 1

    fmt.Println(nextInt()) // Output: 4, because nextInt and newInt2 are different closures with their own state.
}