package main

import (
	"fmt"
)

func main() {
	//Calculate Simple interest from Principal (P), Rate of Interest (R), Time (T)
	//SI = P*R*T

	var principal float64 = 10000
	var roi float64 = 10.5
	var time float64 = 30

	var si float64 = 0

	si = principal * roi * time
	fmt.Printf("Simple interest = is %f\n", si)

}
