package main

import "fmt"

//    1
//   121
//  12321
// 1234321

func main() {
	var n int = 4
	for i := 1; i <= n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}

		for j := i - 1; j > 0; j-- {
			fmt.Printf("%v", j)
		}

		fmt.Println()
	}
}
