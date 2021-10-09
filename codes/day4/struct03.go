package main

import "fmt"

func main() {
	s1 := Studens{"小三", 18}
	fmt.Println(s1.name, s1.age)

	s2 := struct {
		name string
		age  int
	}{
		"阿花",
		16,
	}
	fmt.Println(s2)
	fmt.Println(s2.name, s2.age)

	s3 := Worker{"老孟", 30}
	fmt.Println(s3)
	fmt.Println(s3.string, s3.int)

}

type Studens struct {
	name string
	age  int
}

type Worker struct {
	string // 匿名字段, 默认以数据类型为匿名字段的名字
	int    //	匿名字段,同一类型只能用一次
	// string 多次使用会报错,重复的名字
}
