package main

import "fmt"

func main() {
	test(true)
	fmt.Println("``````````````````")
	test(false)
}

func test(stop bool) {
	fmt.Println("A")
	if stop {
		fmt.Println("B")
		return
	}
	fmt.Println("C")
}
