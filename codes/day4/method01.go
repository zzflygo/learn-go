package main

import "fmt"

type Worker struct {
	name string
	age  int
	sex  string
}

func (w Worker) work() {
	fmt.Println("工人正在工作.....")
}

type Cat struct {
	name string
	age  int
}

func main() {
	w1 := Worker{"老李", 30, "男"}
	w1.work()

	w2 := &Worker{"小花", 28, "女"} //此时w2类型为*main.Worker

	w2.work()    // 函数调用 可以省略指针
	(*w2).work() //不省略指针
	c1 := Cat{"三花", 4}
	w2.printInfo() //两个调用的方法名 一样
	c1.printInfo() //因为接受者不一样所以能够使用
}

func (p Worker) printInfo() { //方法中 函数名可以一样,但是接受者不能一样
	fmt.Printf("工人名字:%s ,工人年龄:%d \n", p.name, p.age)
}
func (p Cat) printInfo() { //方法可以接受者 可以是参数,也可以是指针
	fmt.Printf("猫咪名字:%s ,猫咪年龄:%d \n", p.name, p.age)
}
