package main

import "fmt"

func main() {

	fmt.Println("Regular for loop:-----------")
	for i := 0; i < 10; i++ {
		fmt.Println("I want to join GOOGLE!")
	}

	fmt.Println("For loop like while loop:---------")
	a := true
	b := 0
	for a {
		if b == 10 {
			a = false
			fmt.Println("Done!!!")
		}
		b++
	}

	fmt.Println("For each loop:------------")
	var x = []int{5, 6, 7}
	for i, v := range x {
		fmt.Println(i, v)
	}
}
