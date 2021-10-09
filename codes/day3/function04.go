package main

import "fmt"

/*  值传递 int,float,array

引用传递 切片 ss1:=make([]int,3,8)
		map  ssr:=make(map(int)string)
		传递的是内存地址
*/

func main() {
	s1 := [4]int{1, 2, 3, 4}  //数组是值传递
	fmt.Println("调用前的数据", s1) // 1 . 2 . 3 . 4
	func1(s1)                 // 100. 2 . 3. 4
	fmt.Println("调用后的数据", s1) // 1 . 2. 3 . 4
	fmt.Println("~~~~~~~~~~~~")

	s2 := []int{1, 2, 3, 4} //切片是引用传递
	fmt.Println("调用前的数据", s2)
	func2(s2)
	fmt.Println("调用后的数据", s2)
	fmt.Println("~~~~~~~~~~~~")
}

func func1(n [4]int) {
	n[0] = 100
	fmt.Println("修改后的数据", n)

}

func func2(s3 []int) {
	s3[0] = 100
	fmt.Println("修改后的数据", s3)
}
