package main

import (
	"fmt"
)

func main() {

	var base int = 0
	var i int = 0

	for base <= 100 {
		i++
		base = 9 * i
		if base >= 100 {
			break
		}
		fmt.Println(base)
	}
	fmt.Printf("个数为%d\n", i-1)
}

//
