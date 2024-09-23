package main

import "fmt"

// Print following pattern
// 1 2 3 4 5
// 1 2 3 4 5
// 1 2 3 4 5
// 1 2 3 4 5
// 1 2 3 4 5

func main() {

	var n int = 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}

}
