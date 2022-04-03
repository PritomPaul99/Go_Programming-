package main

import "fmt"

//import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	arr2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9},}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr2[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("Length of arr2:", len(arr2))

	arr3 := [5] int {1,2,3,4,5}

	for i := 0; i < 5; i++ {
		fmt.Print(arr3[i], " ")
	}
	fmt.Println()

	fmt.Println(len(arr3))
}
