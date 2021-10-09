package main

import "fmt"

func main() {
	//字符串是一些字符的集合,每个字符都有固定的位置(索引,下标,index(索引):从0开始,到长度-1)
	s1 := "hello中国"
	s2 := `hello world`

	fmt.Println(s1)
	fmt.Println(s2)
	//字符串的长度:返回的是字节的个数
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	//获取某个字节
	fmt.Println(s1[0])
	a := `h`
	b := 104
	fmt.Printf("%c, %v, %c \n", s2[0], a, b)

	//遍历 有中文的字符串 用range遍历  用for循环会出现乱码.原因是中文占3个字节

	for i, v := range s1 {
		fmt.Printf("%d %c \n", i, v)
	}
	// 字符串不能更改内容
}
