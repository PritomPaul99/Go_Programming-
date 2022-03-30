package main

import "fmt"

//import "fmt"

func main() {
	var x, z rune
	var y string = "Pritom"

	// var a, b = 10, 20

	x = 'a'

	fmt.Println("x =", x)
	fmt.Print("x =",x, "\n")
	fmt.Printf("x = %c\n", x)

	fmt.Println("y =", y)
	fmt.Printf("y = %s\n", y)

	fmt.Println("x + y =",string(x)+y)

	fmt.Printf("z = %c", z)
}

/*
* % doesn't work on float type
* == / != doesn't work on float type
* rune is a allies of int32





*/