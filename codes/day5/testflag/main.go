package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//flag包返回的是一个指针
	// name := flag.String("name", "小张", "请输入名字")
	// age := flag.Int("age", 9000, "请输入年龄")
	// married := flag.Bool("married", false, "婚姻状态")
	// cTime := flag.Duration("cTime", 0, "请输入时间")

	// flag.Parse() //输入后一定要解析

	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*cTime)

	var (
		n1    string
		a1    int
		m1    bool
		delay time.Duration
	)
	flag.StringVar(&n1, "name", "小灰", "输入名字")
	flag.IntVar(&a1, "age", 800, "年龄")
	flag.BoolVar(&m1, "married", false, "结婚了吗")
	flag.DurationVar(&delay, "多久", 0, "ns开始")
	flag.Parse() //一个函数能只能出现一个解析
	//先解析再使用,不然不能接受命令行输入的内容,打印的为默认.
	fmt.Println(n1)
	fmt.Println(a1)
	fmt.Println(m1)
	fmt.Println(delay)

	fmt.Println(flag.Args()) //testflag.exe a b c d 得到[a b c d]的切片
	//和os.Args[1:]是一样的功能,但是第一个元素不是地址
	fmt.Println(flag.NArg())  //统计flag.Args()输入了多少个元素
	fmt.Println(flag.NFlag()) //统计有几个命令行参数
	//testflag.exe -name 小鸡鸡 -age 16 ad 100 小妞   命令行参数有2个  非命令行参数有[ad 100 小妞]3个
}
