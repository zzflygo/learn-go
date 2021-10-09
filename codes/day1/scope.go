package main

import "fmt"

var name string = "kk"

func main() {
	fmt.Println(name)
	//作用域。 说明标识符的使用范围
	var age int = 31
	var name string = "silence"
	{
		var age int = 33
		fmt.Println(name, age)
		name = "kk"
	}
	fmt.Println(name, age)
}
