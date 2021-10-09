package main

import "fmt"

/*
高阶函数:
	根据go语言的数据类型的特点,可以将一个函数作为另一个函数的参数.
	fun1(), fun2()
	将fun1函数最为fun2这个函数的参数.
		fun2函数:就叫做高阶函数
			接受了一个函数作为参数的函数,叫做高阶函数.
		fun1函数:回调函数
			作为另一个函数的参数的函数,叫做回调函数.
*/

func main() {
	res1 := add
	fmt.Println(res1(10, 20))

	res2 := oper(10, 33, add)
	fmt.Println(res2)

	res3 := oper(10, 20, sub)
	fmt.Println(res3)

	fun1 := func(a, b int) int {
		return a * b
	}
	fmt.Println(fun1(10, 15))

	res4 := oper(14, 10, fun1)
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(res5)

}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
