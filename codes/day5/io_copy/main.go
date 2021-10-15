package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func txtcopy(s, d string) {
	//打开要复制的原文件
	src, err := os.Open(s)
	if err != nil {
		fmt.Println("open src file failed err:", err)
		return
	}
	//关闭原文件
	defer src.Close()

	//打开要copy的目的文件
	dst, err := os.OpenFile(d, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open src file failed err:", err)
		return
	}

	//关闭复制文件
	defer dst.Close()
	//进行copy操作
	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("copy file failed err:", err)
		return
	}
	fmt.Println("copy success")
	fmt.Println(n)

}

func main() {
	a := "./t2.txt"
	b := "./rn.txt"
	txtcopy(a, b)
	n, err := ioutil.ReadFile(b)
	if err != nil {
		fmt.Println(" file open failed err:", err)
		return
	}
	fmt.Println(string(n))

}
