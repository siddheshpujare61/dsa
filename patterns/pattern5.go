package main

import "fmt"

// Print following pattern
// 1
// 1 2
// 1 2 3
// 1 2 3 4
// 1 2 3 4 5

func main() {

	var n int = 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}

}
