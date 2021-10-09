package main

import "fmt"

func main() {
	//冒泡程序

	arr := [5]int{15, 23, 8, 10, 7}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)

	}
}
