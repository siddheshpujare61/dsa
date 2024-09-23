package main

import "fmt"

// Print following pattern
// A
// A B
// A B C
// A B C D
// A B C D E
// A B C D E F

func main() {

	var n int = 70

	for i := 65; i <= n; i++ {
		for j := i; j <= i; j++ {
			fmt.Printf("%c ", j)

		}
		fmt.Println()
	}
}
