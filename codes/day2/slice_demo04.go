package main

import "fmt"

/*
    深拷贝: 值类型的拷贝默认都是深拷贝.

	浅拷贝: 拷贝的是数据 地址
	导致多个变量指向同一内存
	引用数据 默认都是浅拷贝:slice, map,
	因为切片是引用类型的数据,直接拷贝的是地址.
*/

func main() {

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0)

	fmt.Println("-------------------------------")
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("-------------------------------")
	//copy 函数 copy(s2,s1)把s1的数据copy到s2

	s3 := []int{7, 8, 9}
	copy(s3, s2) //把s2中数据按位复制(深拷贝)到s3 前面为目标切片 后面为来源切片
	fmt.Println(s2)
	fmt.Println(s3) // 如果目标cap容量不够,有多少装多少.
	// 按位取 可以用下标取
	//copy(s2[2:],s3)
	//copy(s2[],s3[1:3])

}
