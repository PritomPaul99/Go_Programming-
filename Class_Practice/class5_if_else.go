package main

import "fmt"

func main() {
	x := 20

	fmt.Println("If-Else:-------------")
	if x%2 == 0 {
		if x%5 == 0 {
			fmt.Println("Divisible by 2 and 5!")
		} else {
			fmt.Println("Divisible by 2, but not 5!")
		}
	} else {
		fmt.Println("Not divisible by 2!")
	}

	println("Ladder:--------------")
	if x%5 == 0 && x%2 == 0 {
		fmt.Println("Divisible by 2 and 5!")
	} else if x%2 == 0 && x%5 != 0 {
		fmt.Println("Divisible by 2, but not  5!")
	} else {
		fmt.Println("Not divisible by 2!")
	}
}
