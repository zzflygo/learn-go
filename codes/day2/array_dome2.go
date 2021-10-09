package main

import "fmt"

func main() {

	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Println(len(a2))
	fmt.Println("一维数组的长度是", len(a2[2]))
	fmt.Printf("二维数组的长度是 %d \n", len(a2))

	for _, x := range a2 {
		for _, y := range x {
			fmt.Println(y)
		}
	}
}
