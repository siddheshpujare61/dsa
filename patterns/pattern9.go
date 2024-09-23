package main

import "fmt"

// 1 1 1 1
// 2 2 2
// 3 3
// 4

func main() {
	var n int = 4
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
	}
}
