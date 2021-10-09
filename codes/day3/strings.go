package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings包的使用.
	s1 := "helloo world zhang"
	fmt.Println("长度", len(s1))
	fmt.Printf("%c \n", s1[5])
	//strings.Contains(s, "x") bool 判断 s字符串里包不包含"x" ,包含为1,不包含为0
	fmt.Println(strings.Contains(s1, "x"))
	//strings.ContainsAny(s,"abcd")判断 s字符串里包不包含"abcd"中的任何一个 ,包含为1,不包含为0
	fmt.Println(strings.ContainsAny(s1, "abcl"))
	//strings.Count 统计字符串abacva中 重复出现几次a
	fmt.Println(strings.Count(s1, "l")) //出现了3次l ,如果没出现则返回0
	fmt.Println(strings.Count("中国人在国外吃中国菜", "国"))
	//strings.Index 返回int值.,子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.Index(s1, "l"))
	fmt.Println(strings.Index(s1, "aad"))
	//strings.IndexAny 子串sep中任意一个元素在s1第一次出现的位置,不存在或者 子串为空 则返回-1。
	fmt.Println(strings.IndexAny(s1, "aad"))
	//strings.Title 把字符串中每个单词的首字母变成大写
	fmt.Println(strings.Title(s1))
	//strings.ToLower 所有字母变小写.数字标点不变
	fmt.Println(strings.ToLower(s1))
	//strings.UoUpper 所有字母变大写.数字标点不变
	fmt.Println(strings.ToUpper(s1))
	//strings.Join 字符串的拼接
	ss1 := []string{"avc", "ccsa", "霓虹", "112"}
	s3 := strings.Join(ss1, "~") //要拼接的切片 用什么拼接   拼接成一个新的字符串
	fmt.Println(s3)
	//strings.Split 把一个字符串切片 成新的切片
	s4 := "123,42123,5331,123,1241"
	ss2 := strings.Split(s4, ",")
	fmt.Println(ss2)
	//strings.Repeat 把一个字符串重复拼接 n 次
	s5 := strings.Repeat(s1, 5) //把hello 拼接5次
	fmt.Println(s5)
	//strings.Replace 替换 把s1中的"l" 替换几次输入,替换所有次数为-1
	s6 := strings.Replace(s1, "l", "^", -3) //把s1中的"l" 替换为"^" 替换2次,负数为全部替换
	fmt.Println(s6)

	//截取子串

	sa := s1[2:15]
	fmt.Println(sa)

}
