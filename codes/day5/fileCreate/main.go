package main

import (
	"fmt"
	"os"
)

func main() {
	fileName1 := "bb.txt"
	// file1, err := os.Create("D:/goland/codes/day5/testflag/aa.txt")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// fmt.Println(file1)
	file2, err := os.Create(fileName1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file2)

	file3, err := os.OpenFile("bb.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(file3)
	file3.Close()
	//删除文件
	//os.Remove 删除文件和空目录
	err = os.Remove("D:/goland/codes/day5/fileCreate/aa.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("Remove Succeed")
}
