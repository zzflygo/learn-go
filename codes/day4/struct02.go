package main

import "fmt"

func main() {
	p1 := Person{"王二狗", 18, "男", "北京"}
	fmt.Printf("内存: %p \n内容: %T\n", &p1, p1)

	p2 := p1
	fmt.Printf("内存: %p \n内容: %T\n", &p2, p2)
	p2.name = "老李"

	//p3 := &p1 简短声明也行
	var p3 *Person //指针要先声明 再 取地址
	p3 = &p1
	fmt.Printf("内存: %p \n内容: %T\n", &p3, p3)
	//(*p3).name = "小狗狗"  可以和数组一样省略*
	p3.name = "小狗狗"
	fmt.Printf("内存: %p \n内容: %v\n", &p1, p1)

	//new() 使用内置函数new()是go语言中专门返回某种类型指针的函数 ,他返回一个指针
	p4 := new(Person)

	fmt.Printf("%T \n", p4)
	fmt.Printf("%p \n", &p4)
	p4.name = "小乌龟"
	p4.age = 30
	p4.sex = "女"
	p4.address = "海底捞"
	fmt.Printf("%p \n", &p4.name)
	p5 := *p4
	fmt.Printf("p5 %T \n", p5)
	fmt.Println(p5)

}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}
