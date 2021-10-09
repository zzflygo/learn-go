package main

import "fmt"

/* 匿名:没有名字
匿名函数:没有名字的函数.

定义一个匿名函数,直接在后面加上()进行调用,通常只能调用一次.
也可以使用匿名函数赋值给某个函数变量,这样就可以多次调用了.

Go语言是支持函数式编程:
	1.将倪明芳
*/

func main() {

	fun1()
	fun2 := fun1
	fun2()

	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()
	fun3()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(10, 20)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数后有括号,代表把函数的返回值赋值给res1
	fmt.Printf("res1的类型:%T,res1的值%v \n", res1, res1)

	res2 := func(a, b int) int {
		return a + b
	} //匿名函数后 没有扩号,代表把函数本省赋值给res2 后面直接调用res2()
	fmt.Println(res2(15, 25))
	fmt.Printf("res2的类型:%T\n", res2)
	fmt.Println(res2)

}

func fun1() {
	fmt.Println("我是一个fun1()函数")
}
