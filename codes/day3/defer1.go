package main

import "fmt"

func main() {
	fmt.Println("start")

	defer func() { //defer 后面必须跟一个函数func() 带括号的, 使用堆栈执行
		fmt.Println("defer1")
	}()
	defer func() {
		fmt.Println("defer2")
	}()
	fmt.Println("end")

}
