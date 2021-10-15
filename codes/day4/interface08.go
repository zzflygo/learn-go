package main

import "fmt"

type A interface{}

type Cat struct {
	name string
}
type Person struct {
	name string
	age  int
}

func test1(a A) {
	fmt.Println(a)
}

func test2(a interface{}) {
	fmt.Println("==>", a)
}

func test3(a []interface{}) {
	for i, v := range a {
		fmt.Printf("第%d个数是->%v \n", i, v)
	}

}
func main() {
	var a1 A = Cat{"小黄"}
	var a2 A = Person{
		name: "张夏",
		age:  16,
	}
	var a3 A = 100
	var a4 A = 3.145
	var a5 A = [...]int{1, 3, 4, 5}
	test1(a1)
	test1(a2)
	test1(a3)
	test1(a4)
	test1(a5)
	test2(a1)
	test2(a2)
	test2(a3)
	test2(a4)
	test2("小老弟")

	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5, "大象", 66)
	test3(slice1)

}
