package main

import "fmt"

func main() {
	var no int = 5
	var sum int = 0
	for i := 1; i <= no; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Println("Sum of odd nos:", sum)
}
