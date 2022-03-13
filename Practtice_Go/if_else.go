package main

import "fmt"

func main() {
	if i := 10; i > 12{
		fmt.Println("10 is greater than 12? True!")
	} else {
		fmt.Println("10 > 12? False!")
	}


	if 5 > 6 {
		fmt.Println("5 > 6")
	}else if 10 > 5{
		fmt.Println("10 > 5")
	}else{
		fmt.Println("!!!!!!")
	}
}