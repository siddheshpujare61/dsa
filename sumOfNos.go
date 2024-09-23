package main

import "fmt"

func main() {
	var no int = 14560
	var sum int = 0

	for i := no; i > 0; i-- {
		sum += no % 10
		no = no / 10
	}
	fmt.Println("Sum of nos:", sum)
}
