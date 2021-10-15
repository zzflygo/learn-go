package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfromFile1() {
	//os.Open 打开文件 要接 file名.Close()关闭
	//因编译器不同.goland不识别相对路径,需要写绝对路径
	txt1, err := os.Open("D:/goland/codes/day5/file_open/main.go") //可以使用绝对路径 和 相对 路径
	defer txt1.Close()                                             //有打开就要有关闭. 用defer 延迟到最后
	if err != nil {
		fmt.Printf("file open failed err:%v", err)
		return
	}
	//读文件 Read函数为 File结构体的一个方法
	//直接用File.Read(长度)来调用,返回个数 和err
	//func (f *File) Read(b []byte) (n int, err error)
	//var tmp = make([]byte,128)  指定读的长度 .常用为1024 或1024的倍数
	var tmp [128]byte // tmp里面存的是数据
	//n为读了多少个字节,结束时返回0和err = io.EOF
	for {
		n, err := txt1.Read(tmp[:]) //可以控制从什么地方开始读写,只能第一次使用 不能放在循环里
		//每次读写的时候都会把读到的内容存储在tmp[128]byte切片里.并记录上次读写位置
		if err == io.EOF {
			fmt.Println("数据已经读完!")
			return
		}
		if err != nil {
			fmt.Println("read file failed err:", err)
			return
		}

		fmt.Println("读了", n, "个字节")
		fmt.Println(string(tmp[0:n]))
	}

}
func readFromFilebyBufio() {
	//打开文件
	txt2, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("err:", err)
	}
	//记得关闭
	defer txt2.Close()
	//创建变量,把读到的东西放在变量里面
	reader := bufio.NewReader(txt2)
	//'\n'表示以换行为基础 按放
	for {
		line, err := reader.ReadString('\n') //注意里面是byte 用单引号' '
		if err == io.EOF {
			fmt.Println("read line has end", err)
			return
		}
		if err != nil {
			fmt.Println("read line failed", err)
			return
		}
		fmt.Print(line) //本来就是按\n来读.用Println的话 会多加一个\n
	}
}
func ioutilRead() {
	//打开一个文件
	txt3, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed", err)
		return
	}
	//记得关闭

	fmt.Println(string(txt3)) //返回的是字节切片,要转成字符串

}
func main() {
	//readFromFilebyBufio()
	ioutilRead()
	s1 := "一直花狸猫"
	err := ioutil.WriteFile("./xx.txt", []byte(s1), 0666) //每次写入都重置一直文件
	if err != nil {
		return
	}

}
