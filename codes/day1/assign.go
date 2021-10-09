package main

import (
	"fmt"
)

func main() {

	var (
		age    = 31
		weight = 138
	)

	fmt.Println(age)

	age = 32

	age, weight = 32, 139 // 赋值 更新变量的值

	fmt.Println(age, weight)

}
