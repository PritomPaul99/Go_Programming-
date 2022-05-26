package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	x := greater(a, b)

	fmt.Println(x, "is greater")
}

func greater(a, b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}
