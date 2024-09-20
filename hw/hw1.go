package main

import "fmt"

// Sum of N nos which are divisbile by 3
func main() {
	var no int = 10
	var sum int = 0
	for i := 1; i <= no; i++ {
		if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println("SUm is ", sum)

}
