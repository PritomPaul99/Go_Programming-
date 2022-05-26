package main

import "fmt"

func main() {
	var arr [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
