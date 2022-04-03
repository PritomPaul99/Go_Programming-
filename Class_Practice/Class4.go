package main

import "fmt"

func main() {
	var a1 = []int{1, 2, 3}
	fmt.Println(a1)
	a1[0] = 10
	fmt.Println(a1)

	var a2 = make([]int, 5, 10)
	fmt.Println(a2)
	fmt.Println("Length of a2:", len(a2))
	fmt.Println("Capacity of a2:", cap(a2))

	a2 = append(a2, 10)
	fmt.Println(a2)
	fmt.Println("Length of a2:", len(a2))
	fmt.Println("Capacity of a2:", cap(a2))

	//Map

	var m1 = map[string]string {"user":"Pritom","Psw":"Pritom Paul"}
	fmt.Println(m1)

	var m2 = map[int]int {0:0, 1:10, 2:20, 3:30}
	fmt.Println(m2[0], m2[1], m2[2], m2[3])

	// var m3 = map[string]string{}//map here

	var m3 = map[string]int{"a":1, "b":2}
	fmt.Println(m3)
	fmt.Println(m3["d"])
	
	x, ok := m3["d"]
	fmt.Println(x, ok)


}
