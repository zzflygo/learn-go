package main

import "fmt"

func main() {
	fmt.Println("老婆的想法: ")
	fmt.Println("买10个包子")

	var text string

	fmt.Print("有卖西瓜的吗:")

	fmt.Scan(&text)

	if text == "y" {
		fmt.Println("买1个西瓜")
	}

}
