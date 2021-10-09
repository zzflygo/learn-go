package main

import "fmt"

/*
按照类型分类
		基础类型 int, float,string,bool
		复合类型 array,slice,map,struct,pointer,function,chan  除了4个基本类型 都是复合类型

按照特点分类
		值类型 int, float,string,bool array
            专递的是数据副本,在更改的时候不会更改原数据.
		引用类型slice
		 	 传递的内存地址,多个变量指向了同一块内存地址.

所以切片是引用类型的数据,储存了底层数组的引用.
*/

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("-----------------1,已有数组直接创建切片--------------")
	s1 := a[:5]  //取1-5的数字. 从下标0-4,所以到下标5,但是不包含下标5
	s2 := a[3:8] //取4-8的数字
	s3 := a[5:]  //6-10, 直接到结尾.结尾可省略,从0开始0可以省略
	s4 := a[:]   //1-10 ,开头结尾都省略,取a数组全部

	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	fmt.Printf("%p \n", &a)
	fmt.Printf("%p \n", s1)

	fmt.Println("-----------------2,长度和内容--------------")
	fmt.Printf("s1 len=%d cap=%d \n", len(s1), cap(s1)) //容量从0到9 就是10
	fmt.Printf("s2 len=%d cap=%d \n", len(s2), cap(s2)) //容量从3到9 就是7
	fmt.Printf("s3 len=%d cap=%d \n", len(s3), cap(s3)) //容量从5到9 就是5
	fmt.Printf("s4 len=%d cap=%d \n", len(s4), cap(s4)) //容量10

	fmt.Println("-----------------3,更改数组内容--------------")
	a[4] = 100 //更改数组 切片会跟着改
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("a的地址:%p\n", &a)
	fmt.Printf("s1的地址:%p\n", s1)

	fmt.Println("-----------------4,更改切片内容--------------")
	s2[2] = 200 //更改切片,数组也会跟着改
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("a的地址:%p\n", &a)
	fmt.Printf("s1的地址:%p\n", s1)

	fmt.Println("-----------------5,追加切片内容-------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println("a:", a) //追加切片数据的时候 底层数组相应位置上的数据会被更改
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Printf("a的地址:%p\n", &a)
	fmt.Printf("s1的地址:%p\n", s1)

	fmt.Println("-----------------5,添加切片扩容-------------")
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 2, 2, 2, 2, 2)
	// 单切片进行扩容时,超出数组范围.会复制原数组数据到新的内存创建新的底层数组
	//此时内存地址会改变,同时操作并不会影响原数组
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)
	fmt.Println(len(s1), cap(s2))
	fmt.Printf("a的地址:%p\n", &a)
	fmt.Printf("s1的地址:%p\n", s1)
}
