package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前a的值:", a)
	fun1(a)
	fmt.Println("函数调用后a的值:", a)

	fun2(&a)
	fmt.Println("fun2函数调用后a的值:", a)

}

//值传递,把值付给了n
func fun1(n [4]int) { //在声明 n 类型时候 会开辟一个内存,把a的值输出给n
	fmt.Println("函数修改前n的值:", n)
	n[0] = 200 //此处修改的是 n 内存里面的值 ,与a里面的无关
	fmt.Println("函数修改后n的值:", n)
}

//引用传递.把a的地址给了n2
func fun2(n2 *[4]int) {
	fmt.Println("函数修改前n的值:", n2)
	(*n2)[0] = 200 // 可以简写为 n2[0] = 200 不用完整的写 .直接修改a里面的值
	fmt.Println("函数修改后n的值:", n2)
}
