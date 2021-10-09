package main

import "fmt"

func main() {
	var name string = `kk`
	var name1 string = "\na\\nb\a"
	fmt.Printf("%s%s", name, name1)
	fmt.Println("my name is " + name)

	//关系预算
	//  > >= < <= != ==
	// s1 s2
	// 字符串最左侧的赐福开始比较
	// abc > abd
	// a b 相同 c > d false
	// abc > abcd 前项空 false

	// =, +=  a+=b  a=a+b
	name += " hi"
	fmt.Println(name)
	//字符串内的数据 进行 索引 长度 切片 必须是ascii byte 会出乱码
	//索引 就是单独提出第几个字符 name[] 中括号
	// name => "kk hi" 0-len(n)-1
	fmt.Printf("%T\n", name[0])
	fmt.Println(name[0])
	// 长度 len（） len（name）
	fmt.Println(len(name))
	//切片 拿出字符串的一部分来
	fmt.Println(name[0:4])

}
