package main

import "fmt"

//数据的长度的固定的 len() cap()都能查看长度
//语法 var name [len]tpye
//     var name = [len]tpye{元素1,元素2....}
//     var name = [...]tpye{元素1,元素2....} 根据元素自定义长度
func main() {
	var i = [5]int{1, 3, 5}
	fmt.Println(i)

	var n [5]int
	fmt.Println(n)

	var y = [5]int{1: 2, 4: 4}
	fmt.Println(y)

	var m = [5]string{2: "王二狗", 4: "ss", 0: "小美"} // 没定义就是空的字符串" "
	fmt.Println(m)

	b := [...]int{1, 5, 7, 8, 9, 22}
	fmt.Println(b)

	p := [...]int{1: 8, 12: 8}
	fmt.Printf("%d", cap(p))
	fmt.Println(p)

}
