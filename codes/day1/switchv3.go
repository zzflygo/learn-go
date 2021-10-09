package main

import "fmt"

//控制台输入一个成绩
//>=90  >>a
//>=80  >>b
//>=70  >>a
//>=60  >>b
//其他  >>e

func main() {
	var score int
	fmt.Print("请输入成绩: ")
	fmt.Scan(&score)

	switch {
	case score >= 90:
		fmt.Println("A")

	case score >= 80:
		fmt.Println("B")

	case score >= 70:
		fmt.Println("C")

	case score >= 60:
		fmt.Println("D")

	default:
		fmt.Println("E")
	}
}
