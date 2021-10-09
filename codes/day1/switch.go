package main

import "fmt"

func main() {
	fmt.Println("老婆的想法")
	fmt.Println("有西瓜卖吗")

	var text string

	fmt.Scan(&text)

	switch text {
	case "y":
		fmt.Println("买一个西瓜")
	}

}
