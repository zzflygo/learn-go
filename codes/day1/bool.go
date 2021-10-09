package main

import "fmt"

func main() {
	isBoby := true
	fmt.Println(isBoby)

	fmt.Printf("%T %t\n", isBoby, isBoby) // %T传递类型 %t数值
	// 操作
	// 逻辑运算
	// and or not
	//与操作 两个操作数 左操作数 && 右操作数
	//运算规则：都为真的结果为真
	//A && B
	fmt.Println("&&", "与预算and")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	//或运算 左操作数||友操作数
	//运算规则：只要一个为真结果未真
	fmt.Println("||", "或运算or")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	//非 ，!操作数
	//运算规则：你真我假 你假我真
	//!A 结果反过来
	fmt.Println("!", "取反")
	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println("==", "等于")
	fmt.Println(true == true)
	fmt.Println(true == false)
	fmt.Println(false == false)

	fmt.Println("!=", "不等于")
	fmt.Println(true != true)
	fmt.Println(true != false)
	fmt.Println(false != false)
	// Printf 占位符的符号 %T传递类型 %t数值
}
