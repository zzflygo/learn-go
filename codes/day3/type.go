package main

import "fmt"

func main() {

	var fs []func(int, int) int //定义fs 为一个切片,存储 func(int,int)int的函数的切片

	fs = append(fs, add, mul) //把 add,mul 函数 储存在fs切片里面

	for _, f := range fs { //遍历fs  等于把fs里面的函数 f:=add  ,f:=mul
		fmt.Printf("f的类型:%T  ", f)
		fmt.Println(f(8, 6)) //此时f的类型应该与add,mul一样.所以使用f(8,6)来调用 并返回值
	}

}

func mul(a, b int) int {
	return a * b
}
func add(a, b int) int {
	return a + b
}
