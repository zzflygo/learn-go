package main

import "fmt"

func main() {
	height := 1.691245
	fmt.Printf("%T,%f,%g,%e", height, height, height, height)
	// %T 类型 %f 小数 %g自动 %e科学计数法
	// 运算
	// 算数运算
	// + - * / ++ --
	fmt.Println(1.2 + 1.1)
	fmt.Println(1.2 - 1.1)
	fmt.Println(1.2 * 1.1)
	fmt.Println(1.2 / 1.1)
	height++
	fmt.Println(height)
	height--
	fmt.Println(height)

	// 关系运算 浮点数在关系运算时没有等于 和不等于 因为不是精确运算
	//> >= <= <

	fmt.Println(1.2 > 1.1)
	fmt.Println(1.2 < 1.1)
	fmt.Println(1.2 >= 1.1)
	fmt.Println(1.2 <= 1.1)
	//赋值运算
	//+=, -=, /= , *= a+=b  a=a+b
	//类型转换 float64() float32()  float转int会被截断
	fmt.Printf("%f %10.3f\n", height, height)
	//%f  %10.3f  10表示 占位10 3表示小数点后3位
}
