package main

import (
	"fmt"
	"sort"
)

func main() {
	a := 4
	fmt.Printf("%T\n", a)
	b := [3]int{1, 2, 3} //数组
	fmt.Printf("%T\n", b)
	c := make([]int, 8, 20) //切片 长度为8 容量为20
	fmt.Printf("%T %d %d\n", c, len(c), cap(c))
	d := make(map[int]string) //映射 键 int  值 string
	d[2] = "gold"
	d[8] = "week"
	fmt.Printf("%T %d \n", c, len(d))    //长度为9
	e := make(map[string]map[int]string) //多为映射,键位string 值为map[int]string
	e["hi"] = map[int]string{1: "ff"}
	e["zz"] = map[int]string{2: "ss"}

	fmt.Printf("%T            长度: %d \n", e, len(e)) //长度为2
	fmt.Println(e)
	fun1()
	fmt.Printf("%T  \n", fun1)

}

func fun1() {
	s1 := []int{8, 11, 85, 77, 32, 65, 8, 99, 112}
	fmt.Println(s1)
	sort.Ints(s1) //冒泡排列 从大到小
	fmt.Println(s1)
}
