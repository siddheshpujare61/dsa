package main

import "fmt"

// 1
// 2 3
// 4 5 6
// 7 8 9 10
// 11 12 13 14 15

func main() {
	var n int = 5
	var num int = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			num += 1
			fmt.Printf("%v ", num)
		}
		fmt.Println()
	}
}
