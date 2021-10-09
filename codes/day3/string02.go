package main

import (
	"fmt"
	"strconv"
)

//go语言中 +号 两边的类型要求一致

func main() {
	// strconv包 字符串 和 基本类型之间的转换
	// strings convert

	// i, err := strconv.Atoi("-42") 表示把string类型转成int数值
	// s := strconv.Itoa(-42) 表示把int转化成string类型

	// strconv.Parse  Parse 解析 string 转化成其他类型. bool float int Uint
	// b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415",64)
	// i, err := strconv.ParseInt("-42",10,64)
	// u, err := strconv.ParseUint("42",10,64)

	//strconv.Format Format 的格式化 想把什么类型格式化为string 用Format

	//把字符串转化为其他类型 用Parse 解析  都会返回一个err
	//把其他类型格式化为字符串 用Formate 没有err

	s1 := "true"

	b1, err := strconv.ParseBool(s1)

	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%t %T \n", b1, b1)

	b2 := strconv.FormatBool(b1)
	fmt.Printf("%s %T \n", b2, b2)

	a1 := strconv.Itoa(-42)
	fmt.Printf("%s %T \n", a1, a1)

	a2, err := strconv.Atoi(a1)
	fmt.Printf("%d %T \n", a2, a2)

	b3, err := strconv.ParseBool("false")
	f1, err := strconv.ParseFloat("3.1412", 64) //把3.1412 按float64进行转换
	i1, err := strconv.ParseInt("1314", 10, 64) //把1314 按10进制 int64进行转换

	fmt.Printf("%t %T \n", b3, b3)
	fmt.Printf("%f %T \n", f1, f1)
	fmt.Printf("%d %T \n", i1, i1)

}
