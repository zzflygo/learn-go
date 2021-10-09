package main

import "fmt"

func main() {
	fmt.Println("老公的想法")
	fmt.Println("有西瓜卖吗")

	var text string

	fmt.Scan(&text)

	switch text {
	case "y":
		fmt.Println("买一个包子")
	default:
		fmt.Println("买10个包子")
	}

}
