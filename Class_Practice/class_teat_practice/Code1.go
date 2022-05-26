package main

import (
	"fmt"
	// "math"
)

func main() {
	var a, b, sum int
	fmt.Scanf("%d %d", &a, &b)
	sum = a + b
	// var multi int = 1
	multi := 1
	for i := 0; i < 3; i++ {
		multi = sum*multi
	}
	fmt.Printf("Sqr of (a+b); %d", multi)
}
