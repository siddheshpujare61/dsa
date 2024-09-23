package main

import "fmt"

// Print following pattern
// 1
// 2 1
// 3 2 1
// 4 3 2 1
// 5 4 3 2 1

func main() {

	var n int = 5

	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}

}
