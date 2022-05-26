package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	p := "Akhi"
	q := "Sarker"

	a, b := swap(p, q)

	fmt.Println(a,b)
}
