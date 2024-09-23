package main

import "fmt"

// Print following pattern
// A
// B C
// D E F
// G H I J

func main() {

	var n int = 68
	var num int = 65
	for i := 65; i <= n; i++ {
		for j := 65; j <= i; j++ {
			fmt.Printf("%c ", num)
			num += 1
		}
		fmt.Println()
	}

}
