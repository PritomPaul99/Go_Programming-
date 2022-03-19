package main

import "fmt"

func main() {
	
	a := 5

	for i := 0; i < 10; i++ {
		a--
		fmt.Println("a =", a)
	}	
}