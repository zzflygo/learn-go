package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入你的名字:")
	fmt.Scan(&name) // 接受输入内容 ，赋值给变量name
	// &name => 取name指针（地址）
	fmt.Println("你出入的内容是:", name)
	fmt.Println("aa", "bb", "cc", name)

	var age int
	fmt.Print("请输入你的年龄:")
	fmt.Scan(&age)
	fmt.Print("你输入的年龄是:", age)
}
