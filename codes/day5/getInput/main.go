package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanInput() {
	var s string
	fmt.Println("请输入内容")
	//使用Scan接受只能接受到第一个空格前面的数据
	fmt.Scanln(&s)
	fmt.Printf("%#v", s)
}

func bufioInput() {
	//os.Stdin为标注输入.还有os.Stdout os.Stderr
	reader := bufio.NewReader(os.Stdin)
	//bufio.NewReader()可以从文件里面读.可以从变量里面读.也可以从标准输入读
	//反之bufio.NewWriter()也可以直接写进标准输入
	fmt.Println("请输入内容")
	s, _ := reader.ReadString('\n')
	fmt.Printf("%#v", s)

}

func main() {
	//scanInput()
	bufioInput()
}
