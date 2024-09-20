package main

import "fmt"

func main() {
	var ch string = "X"
	if ch >= "a" && ch <= "z" {
		fmt.Println("Character is lowercase")
	}
	if ch >= "A" && ch <= "Z" {
		fmt.Println("Character is Uppercase")
	}
}
