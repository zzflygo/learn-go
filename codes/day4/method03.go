package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	v := []int{1, 2, 3} //range 在执行时会先确定切片长度..在确定循环的次数
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
