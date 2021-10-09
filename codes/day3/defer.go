package main

import "fmt"

// 当一个函数中有多个延迟调用时,他们被添加到一个堆栈中,
//并在Last in First Out(LIFO)后进先出的顺序中执行
//defer函数在调用的时候.就已经传递了参数.只是最后才打印
//在一个函数中有return时,在return之前,所有的defer 函数先完成 再return.最后外围函数才结束
func main() {
	fmt.Println("~~~~~~~~~~")
	n := 2
	defer fun2(n) //根据栈原则:先进后出 后进先出
	n++
	fmt.Println(n)

}

func fun1(i string) {
	fmt.Println(i)
}

func fun2(n int) {
	fmt.Println("fun2中的数据", n)
}
