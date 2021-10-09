package main

import "fmt"

func main() {
	// 定义变量names位元素类型为字符串长度为55的数组
	var names [55]string //string 55人

	var scores = [...]int{3: 100, 1: 88}
	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)
	fmt.Printf("%T\n", scores)

	fmt.Println(scores)
	// 零值
	// n个对应元素类型的零值组成的宿主

	// 索引 0 >> len-1  0 1 2 3 4
	//[]type{v1,v2,v3,.....,vn} var scores [5]int= [5]int{1,2,3,4,5}
	//[]type{i1:v1,vi:vi,vi:vn} var scores [5]int= [5]int{3: 100, 1: 88}
	//[...]type{v1:v1,ii:vi,in:vn} 最大的索引+1 var scores = [...]int{3: 100, 1: 88}
	//[...]type{v1,v2,v2,...,vn}  元素的数量

	//操作
	//关系运算 != 和 ==
	var nums [4]int = [...]int{100, 88, 0, 20}
	fmt.Println(nums == scores)
	//访问 和 修改
	// 索引
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	nums[0] = 101
	nums[1] = 888

	fmt.Println(nums[0])
	fmt.Println(nums)

	//如何计算数组长度
	fmt.Println(len(nums))

	//遍历
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	for i, v := range nums {
		fmt.Println(i, v)
	}
}
