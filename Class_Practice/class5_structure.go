package main

import "fmt"

func main() {
	type student struct {
		firstName  string
		familyName string
		id         int
	}
	x := student{firstName: "Pritom", familyName: "Paul", id: 10}

	fmt.Println(x)
}
