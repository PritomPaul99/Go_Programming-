package main

import "fmt"

func main() {
	var arr [10]int

	n := len(arr)

	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%d", &arr[i])
	// }
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
