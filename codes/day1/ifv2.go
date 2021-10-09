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

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
