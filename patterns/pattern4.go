package main

import "fmt"

// // Print following pattern
// 1
// 2 2
// 3 3 3
// 4 4 4 4
// 5 5 5 5 5

func main() {

	var n int = 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}

}
