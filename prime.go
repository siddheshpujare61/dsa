package main

import "fmt"

func main() {

	var n int = 3
	//fmt.Println(checkPrime(n))
	var isPrime bool = true
	for i := 2; i*i <= n; i++ { //checking i value upto root of n
		if n%2 == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Prime :", isPrime)

}

// func checkPrime(n int) string {

// 	if n == 1 {
// 		return "prime"
// 	}
// 	for i := 2; i <= n-1; i++ {
// 		if n%i == 0 {
// 			return "not prime"
// 		}
// 	}
// 	return "prime"
// }
