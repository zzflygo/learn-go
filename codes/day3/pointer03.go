package main

import "fmt"

/*
	函数指针:是一个指针,指向了一个函数的指针.
		go语言中, function()函数,默认看做一个指针,前面没有*

		slice ,map ,function 都是可以默认看做没有*的指针

	指针函数:是一个函数, 该函数返回值是一个指针.


*/

func main() {
	var a func()
	a = fun1
	a()

	arr1 := fun2()
	fmt.Println(arr1)

	var arr2 *[4]int //存储了4个指针地址 数组的指针
	arr2 = fun3()
	fmt.Printf("arr2的类型:%T ,arr2的地址%p,arr2的内容%p \n", arr2, &arr2, arr2)

}

func fun1() {
	fmt.Println("fun1........")
}

func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("%p \n", &arr)
	return &arr
}
