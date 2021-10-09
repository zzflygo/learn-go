package main

import "fmt"

func main() {
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 18
	p1.sex = "男"
	p1.address = "北京市"
	fmt.Println(p1)

	p2 := Person{}
	p2.name = "李小花"
	p2.age = 18
	p2.sex = "女"
	p2.address = "昆明市"
	fmt.Printf("%#v\n", p2)

	p3 := Person{name: "张三", age: 30, sex: "男", address: "玉溪市"}
	fmt.Println(p3)

	p4 := Person{
		name:    "老亮",
		age:     33,
		sex:     "男",
		address: "弥勒",
	}
	fmt.Println(p4)

	p5 := Person{"小灰灰", 37, "男", "南宁市"}
	fmt.Println(p5)

	fmt.Println(&p1.name)
	fmt.Println(&p1.age)
	fmt.Println(&p1.sex)
	fmt.Println(&p1.address)

}

type Person struct { //命名大小写 是能不能被别的包引用
	name         string
	age          int
	sex, address string // 同类型的可以,号隔开
}
