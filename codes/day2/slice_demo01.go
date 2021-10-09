package main

import (
	"fmt"
)

//切片 同数组类似,也叫做动态数组 或者变长数组
//		是一个引用类型的容器,指向了一个底层数组

//make()
//	func make(t Type, size...IntegerType) type

//第一个参数:类型
// 		slice, map, chan
//第二个参数:长度len
//  	实际存储元素的数量
//第三个参数:容量的cap
//		最多能够存储的元素数量

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4}
	fmt.Println(s2)
	fmt.Printf("%T, %T \n", arr, s2)

	s3 := make([]int, 3, 8)
	fmt.Println(s3)
	fmt.Printf("容量:%d ,长度:%d \n", cap(s3), len(s3)) //超过容量会自动扩容
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	// fmt.Println(s3[3])  // panic: runtime error: index out of range [3] with length 3

	//append()  切片专用追加数据操作
	s4 := make([]int, 0, 5)
	fmt.Println(s4)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)
	s4 = append(s4, 3, 4, 5, 6, 7) //已经超出容量的cap 但是会自动扩容 成倍数增长
	//切片一旦扩容,就是重新指向一个行的底层数组
	fmt.Println(s4)

	s4 = append(s4, s3...) //再追加其他数组进来要结尾添加...
	fmt.Println(s4)

	//遍历切片
	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}
	for i, v := range s4 {
		fmt.Printf("%d--->%d \n", i, v)

	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	sc := []string{"阿卡", "被他", "诺诺", "小狗", "face", "大头", "小弟", "卡尔", "斯文"}
	fmt.Println(sc)
	sc = sc[0 : len(sc)-1] // 去掉最后一位len()-1
	fmt.Printf("%q \n", sc)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	//删除3
	//copy(dst,src)  用code的办法删除中间位数
	copy(sc[3:], sc[4:])
	fmt.Printf("%d  %d \n", len(sc), cap(sc))
	sc = sc[:len(sc)-1]
	fmt.Println(sc)
	fmt.Printf("%d  %d \n", len(sc), cap(sc))

	for i, v := range sc {
		fmt.Printf("%d-->%v \n", i, v)
	}

}
