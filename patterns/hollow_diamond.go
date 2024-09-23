package main

import "fmt"

//    1
//   121
//  12321
// 1234321

func main() {
	var n int = 4
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		if i != 0 {
			for j := 0; j < 2*i-1; j++ {
				fmt.Printf(" ")
			}
			fmt.Printf("*")
		}

		fmt.Println()
	}
	for i := 0; i < n-1; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		if i != n-2 {
			for j := 0; j < 2*(n-i)-5; j++ { // 2(n-2-i) - 1 -> 2n-4-2i-1 -> 2n-2i-5 -> 2(n-i)-5
				fmt.Printf(" ")
			}
			fmt.Printf("*")
		}

		fmt.Println()
	}
}
