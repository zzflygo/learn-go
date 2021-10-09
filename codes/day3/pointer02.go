package main

import "fmt"

/*
	数组指针 :首先是一个指针, 储存一个数组的地址
	指针数组 :首先是一个数组,储存的数据类型是指针

	*[5]float64   指针, 存储一个有5个浮点数的数组的地址
	*[3]string    指针, 存储一个有3个字符串的数组的地址
	[3]*string	  数组,	存储了3个字符串地址的指针的数组
	[5]*float64	  数组, 存储了5个浮点数地址的指针的数组
	*[5]*float64  指针, 存储了一个有5个浮点数地址的指针的 数组的地址
	*[3]*string	  指针, 储存了一个有3个字符串地址的指针的 数组的地址
	**[4]string   指针, 一个储存了4个字符串地址的 指针 的指针
	**[4]*string  指针, 一个存储了4个字符串地址的数组的指针 的指针




*/

func main() {

	//1. 创建一个普通数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	//1,创建一个数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("arr1的地址%p \n", &arr1)
	fmt.Printf("p1的地址%p \n", &p1)

	//3. 根据数组指针, 操作数组
	(*p1)[0] = 200 //(*p1)[0] = 100 原本写法.可以简写为p1[0]
	fmt.Println(arr1)
	p1[0] = 100
	fmt.Println(arr1)

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}      // a , b, c, d的值 复制给aar2
	arr3 := [4]*int{&a, &b, &c, &d} //a,b,c,d的地址给aar3
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	*arr3[0] = 200
	fmt.Println(a)
	b = 999
	fmt.Println(*arr3[1])

	for i, _ := range arr3 {
		fmt.Printf("%d, %d \n", i, *arr3[i])
	}

}
