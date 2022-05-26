package main

import "fmt"

//import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("3rd angle: %d", (180-(a+b)))
}