package main

import "fmt"

//引出接口实例
//定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak()方法的变量 都是speaker的类型.方法签名
}
type cat struct {
	name string
}
type dog struct {
	name string
}
type peple struct {
	name string
}

func (c cat) speak() {
	fmt.Println(c.name, "喵喵喵")
}
func (d dog) speak() {
	fmt.Println(d.name, "汪汪汪")
}
func (p peple) speak() {
	fmt.Println(p.name, "哎呀")
}

//接受一个参数.传进来什么 就打什么
func da(x speaker) {
	x.speak() //挨了打就要叫
}

func main() {
	c1 := cat{"小花"}
	var d1 = dog{"大黄"}
	var p1 peple
	p1.name = "小灰灰"

	c1.speak()
	d1.speak()
	p1.speak()

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
