package main

import "fmt"

//range 是范围的意思.
//不需要操作数组下标,到达数组的末尾,自动结束for range 循环.
// 每次都从数组中获取 下标和对应的数值.
func main() {

	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i*2 + 1
	}
	fmt.Println(arr1)

	for i, v := range arr1 {
		fmt.Printf("标号是:%d.数据是:%d\n", i, v)
	}
	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println(sum)
}
