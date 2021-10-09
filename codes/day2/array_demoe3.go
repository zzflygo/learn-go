package main

import "fmt"

func main() {
	slice := []int32{1, 2, 3, 4, 5, 6}

	slice2 := slice[:2]
	slice3 := slice[:4]
	fmt.Printf("slice2容量：%v ,长度%v\n", cap(slice3), len(slice3))
	fmt.Printf("slice2容量：%v\n", cap(slice))

	slice2 = append(slice2, 50, 60, 70, 80, 90)
	fmt.Printf("slice为：%v\n", slice)
	fmt.Printf("操作的切片：%v\n", slice2)
	_ = append(slice2, 50, 60)
	fmt.Printf("slice为：%v\n", slice)
	fmt.Printf("操作的切片：%v\n", slice2)
}
