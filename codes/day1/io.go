package main

import (
	"fmt"
)

func main() {
	fmt.Println("1.我叫kk") // Println 打印会自动添加换行
	fmt.Println("2.我叫kk")
	fmt.Print("3.我叫kk") // Print 打印不会加换行
	fmt.Print("3.我叫kk")

	name := "kk"
	fmt.Printf("5.我叫%s\n", name) // \n 换行  字符串的占位%s
	fmt.Printf("6.我叫%s", name)   // Printf 通过占位定义标量填充
}
