package main

// slices package provides functions for working with slices which are more flexible and powerful than arrays.
import (
	"fmt"
	"slices"
)
func main() {
	// Arrays in Go are fixed-size collections of elements of the same type. They have a specific length that is determined at the time of declaration and cannot be changed.
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("Array:", arr)
	fmt.Println("array length", len(arr))

		// declare and initializa an array into a single line.
	  arrayB := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", arrayB)

	// Slices in Go are dynamic, flexible views into the elements of an array. They are more commonly used than arrays because they can grow and shrink in size.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	fmt.Println("slice length", len(slice))
 	var s []string
  fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = append(s, "hello")
	fmt.Println("append:", s)

	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

		t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
}