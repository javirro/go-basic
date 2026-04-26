package main

import (
	"fmt"
)


func main() {

  i:= 3
	for i < 10 {
		fmt.Println("i:", i)
		i++
	}

	for j:= 0; j < 5; j++ {
		fmt.Println("j:", j)
	}

    for i := range 3 {
        fmt.Println("range", i)
    }

    for n := range 6 {
        if n%2 == 0 {
					fmt.Printf("skipping loop")
            continue
        }
        fmt.Println(n)
    }

		var verify bool = true
		if verify {
			fmt.Println("verify is true")
		} else {
			fmt.Println("verify is false")
		}

		var age int = 18
		if age < 18 {
			fmt.Println("You are a minor.")
		} else if age == 18 {
			fmt.Println("You just became an adult.")
		} else {
			fmt.Println("You are an adult.")
		}
}
