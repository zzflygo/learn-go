package main

import "fmt"

func main() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1 //跳过了2-2 2-3 2-4 直接 3-1 3-2开始
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
