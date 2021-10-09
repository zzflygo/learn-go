package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		   map的遍历:
		   		使用: for range
				   数组, 切片: inde, value
				   map:key, value

	*/

	var map1 = make(map[int]string)
	map1[1] = "红孩儿"
	map1[2] = "小钻风"
	map1[3] = "白骨精"
	map1[4] = "白素贞"
	map1[5] = "金角大王"
	map1[6] = "王二狗"
	//1. 遍历map  map是无序的.

	for k, v := range map1 { //map是无序的.再range的时候会随机排列
		fmt.Println(k, v)
	}

	fmt.Println("-------------------------------")
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, map1[i])
	}
	/*  1.获取所有的key,--->切片 / 数组
		3.进行排序.
	    2.遍历key, ---->map[key]
	*/

	keys := make([]int, 0, len(map1))
	fmt.Println(keys)

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	//冒泡排序, sort包排序的. sort.Ints()  sort.Strings
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, map1[key])
	}

	s1 := []string{"Apple", "Windos", "Orange", "abc", "acc", "王二狗", "acd"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)

}
