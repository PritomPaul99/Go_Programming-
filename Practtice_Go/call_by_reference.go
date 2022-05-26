package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	n := len(a)
	fmt.Println("greatest value is:", greater(a, n))
}

func greater(arr []int, n int) int {
	var gr int
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			gr = arr[i]
		}else
		{
			gr = arr[i+1]
		}
	}
	return gr
}
