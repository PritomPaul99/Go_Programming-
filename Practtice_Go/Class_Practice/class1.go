package main

import "fmt"

func main() {
	// a := 10
	// var b int = 20

	// fmt.Println("Add:", a+b)
	// fmt.Println("Sub:", a-b)
	// fmt.Println("Multi:", a*b)
	// fmt.Println("Div:", float64(a)/float64(b))

	var x int8 = 10
	var y int16 = 20

	z := x + int8(y)

	fmt.Println("x + y =", z)
}

/*
int
float
Rune & Strings
complex
Bool..........ei gula sikhbo

int8
int16
int32 -> int
int64 -> int
uint8 -> byte
uint16
uint32 -> uint
uint64 -> uint

*byte is allies for uint8
*Compile time vs Run time


*/
