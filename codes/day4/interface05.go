package main

import "fmt"

//空接口 没有必要取名字 不用声明
//interface{}
//所有的类型都实现了空接口..任意类型的变量都能保存为空接口类型
func shou(a interface{}) {
	fmt.Printf("%T-->%v \n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "小灰"
	m1["age"] = 16
	m1["address"] = "广西"
	m1["married"] = false
	m1["skill"] = [...]string{"rap", "唱", "跳"}
	fmt.Println(m1)
	fmt.Printf("类型:%T \n", m1)
	shou("大人")
	shou(11)
	shou(`乱按`)
	shou(m1)
}
