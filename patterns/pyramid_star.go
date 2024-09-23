package main

import "fmt"

//    *
//   ***
//  *****
// *******

func main() {
	var n int = 4
	var char string = "*"
	for i := 1; i <= n; i++ {

		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Printf("%v", char)
		}

		for j := i - 1; j > 0; j-- {
			fmt.Printf("%v", char)
		}

		fmt.Println()
	}
}
