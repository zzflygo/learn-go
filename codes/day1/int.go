package main

import "fmt"

func main() {
	var age int = 16
	var char byte = 'a'      // 97 => ascii编码 a => 整数值
	var codepoint rune = '我' // unicode编码 我 =>整数值
	fmt.Println(age, char, codepoint)
	fmt.Printf("%T %T %T\n", age, char, codepoint)

	// 算数运算
	// + - * / %
	fmt.Println(1 + 2) // 3
	fmt.Println(1 - 2) // -1
	fmt.Println(1 * 2) // 2
	fmt.Println(1 / 2) // 0  int/int =int 。向下取整
	fmt.Println(1 % 2) // 1 取余数 前小 余数未前数
	//自增 ++ ，自减 --
	// age ++ = age+1     age-- age = age -1
	// 关系运算  => bool
	// > >= < <= != ==
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)

	// 位运算
	// & | ^ << >> &^
	// 两个整数转为二进制进行计算 对应的位进行计算
	// & 与 两个为1 => 1
	// | 或 只要有一个为1 => 1
	// ^ 异或  相同为0 不同为1
	// << 左移 二进制 移位 2^n 乘2的n次方
	// >> 右移    /2^n 除以2的n次方
	//&^ a &^ b  将运算符左边数据相异的位保留，相同位清零。

	//赋值
	// = += -+ *= /= %= &= |= ^= <<= >>=
	//a += b  a= a+b
	//a -= b  a= a-b
	a := 10 //int int8 uint8
	a += 5
	a *= 3
	fmt.Println(a)
	//不同数据类型之间不能进行计算
	var b int8 = 5
	fmt.Println(a + int(b))
	fmt.Println(int8(a) + b)
	// int 64 int8 -128-127

	a = 128
	fmt.Printf("%10d %d\n", a, a)
	fmt.Printf("%010d %-10d %010d %d\n", a, a, a, a)
}
