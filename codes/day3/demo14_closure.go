package main

import "fmt"

/* 闭包(closure)
一个外层函数中,有内层函数,该内层函数中,
会操作外层函数的局部变量(外层函数中的参数,或者外层函数中直接定义的变量),
并且改外层函数的返回值就是这个内层函数.

这个内层函数 和 外层函数的局部变量,统称为闭包结构.

局部标量的生命周期会发证改变:
	正常的局部变量:随着函数调用而创建,随着函数的结束而销毁.
	但是闭包结构中的外层函数的局部变量,
				并不会随着外层的函数的结束而销毁,因为内层函数还要继续使用.
*/

func main() {
	res1 := increment() // 只要出现()就相当于调用了...函数内部的代码 会全部执行一次
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println(res2())
	fmt.Println(res2())
	fmt.Println(res1())

}

func increment() func() int { //外层函数
	//1.定义一个局部变量
	i := 0
	//2.定义一个匿名函数,并给变量自增冰返回
	fun := func() int { //内层函数
		i++
		return i
	}

	//3.返回该匿名函数
	return fun
}
