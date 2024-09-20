package main

import "fmt"

func main() {
	var n int = 5
	factorial := 0

	for i := 1; i <= n; i++ {
		factorial += i
	}
	fmt.Println("factorial is", factorial)
}
