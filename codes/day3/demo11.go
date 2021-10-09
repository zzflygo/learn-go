package main

import "fmt"

/*
GO语言的数据类型:
	数值类型:整数,	浮点
		进行运算操作, 加减乘除,打印
	字符串:
		可以获取单个字符,截取子串,遍历,stings函数的操作
	数组,切片,map..
		存储数组, 修改数组,获取数据,遍历数据..
	函数:
		加(),进行调用

		注意点:函数座位一个复合数据类型,可以看做是一种特殊的变量.
			函数名():将函数进行调用,函数的代码会全部执行,然后将return的结果返回给调用处
			函数名:指向函数体的内存地址

*/

func main() {
	fmt.Printf("%T \n", fun1)
	fmt.Println(fun1)

	var c1 func(int, int)
	fmt.Println(c1)

	c1 = fun1
	fmt.Println(c1)
	c1(100, 200)
	fun1(12, 16)
	res1 := fun2
	res2 := fun2(10, 20)
	fmt.Println(res1(5, 7))
	fmt.Println(res2)

}

func fun1(a, b int) {
	fmt.Printf("a:%d, b:%d \n", a, b)
}

func fun2(a, b int) int {
	return a + b

}
