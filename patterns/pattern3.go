package main

import "fmt"

// Print following pattern
// 1 2 3
// 4 5 6
// 7 8 9

func main() {

	var n int = 3
	var num int = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%v ", num)
			num += 1
		}
		fmt.Println()
	}

}
