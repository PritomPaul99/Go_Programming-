//Variable shadawing
package main

import "fmt"

//import "fmt"

func main() {
	x := 10

	fmt.Println("Shadowing:-")
	if x > 5 {
		fmt.Println(x)

		x := 5 //shadowing is not subjected, Shadawing
		fmt.Println(x)
	}
	fmt.Println(x) //Shadowed


	fmt.Println("Without shadowing:-")
	if x>5 {
		fmt.Println(x)

		x = 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
