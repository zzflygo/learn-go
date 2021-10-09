package main

import "fmt"

/* 指针是指储存一个变量的内存地址的变量,叫做指针变量
&是取地址符 *是取对应地址的数据的符号
*/

func main() {
	a := 10

	var p1 *int //不能简短声明 p1 := *int 是不能使用的
	// *int 指针类型 和变量类型对应 *int *string *array *slice 也有**int表示指针类型的指针
	p1 = &a // &为取地址符号
	fmt.Println(a)
	fmt.Println(p1)
	fmt.Printf("p1的类型:%T \n", p1)
	fmt.Println(*p1) // * 为取储存地址对应的值的符号

	a = 20 //对原数据进行更改 指针取值也会更改
	fmt.Println(a)
	fmt.Println(*p1)

	*p1 = 1111 //使用指针操作原数据 ,原数据也会更改
	fmt.Println(a)
	fmt.Println(*p1)

	var p2 **int    //声明 p2 为*int类型指针的指针
	p2 = &p1        //把p1指针的地址赋予给p2
	fmt.Println(p2) //此时p2里面的内容为p1的地址
	fmt.Printf("p1的地址:%p \n", &p1)
	fmt.Printf("a的地址:%p \n", &a)
	fmt.Printf("*P2的数据:%v \n", *p2) // *p2 为取p2存储的地址的对应的数据.
	// p2 存储的为p1的地址..p1地址存储为a的地址.
	fmt.Printf("**P2的数据:%v \n", **p2)
	fmt.Printf("a的值:%d \n", a)

	//所以*p2为取p2存的p1的地址的值  .*p2 = a的地址.
	//**p2 为取p2存的p1的地址,a的地址的值 **p2 =a的值 111

}
