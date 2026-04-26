package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

// In Go, you can omit the type of the parameters if they are the same. So, the above function can be rewritten as:
func plusPlus(a, b, c int) int {
    return a + b + c
}

// multiple return values
func vals() (int, int) {
    return 3, 7
}

// variable input parameters
func sum(nums ...int) int {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
		return total
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

		a, b := vals()
		fmt.Println("vals() returns:", a, b)
		total := sum(1, 2, 3, 4, 5)
		fmt.Println("Total sum:", total)
		total = sum(10, 20)
		fmt.Println("Total sum:", total)
}