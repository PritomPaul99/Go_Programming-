package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	n := len(arr)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// for i := 0; i < n; i++ {
	// 	fmt.Print(arr[i], " ")
	// }
	fmt.Println(arr)
}
