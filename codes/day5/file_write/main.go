package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"writebyBufio"
)

func filewrite1() {
	//打开文件 没有就创建                         os.O_APPEND会在上一个数据最后添加数据
	file1, err := os.OpenFile("./t1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed err:", err)
	}
	//关闭文件
	defer file1.Close()
	note := "希望的田野上\n"
	file1.WriteString(note)

}
func fileWrite2() {
	file2, err := os.OpenFile("t2.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open file failed err:", err)
	}
	//关闭文件
	defer file2.Close()
	writer := bufio.NewWriter(file2)
	for i := 0; i < 3; i++ {
		n := fmt.Sprintf("小狗汪汪队%d~\n", i+1)
		writer.Write([]byte(n)) //先讲数据写入缓存
		//write.WriteString(n)
	}
	writer.Flush() //@@@重要将缓存的内容写入文件
}

//ioutil.WriteFile自动创建/打开文件 自动关闭 并重置文件.写入
func ioutilWrite() {
	n := "沙县小王子"
	err := ioutil.WriteFile("./t3.txt", []byte(n), 0666)
	if err != nil {
		fmt.Println("open file failed err:", err)
	}

}

func main() {
	// filewrite1()
	// fileWrite2()
	// ioutilWrite()
	writebyBufio.Add("我是一个小花花")

}
