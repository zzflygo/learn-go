package main

import (
	"fmt"
	//同一个父目录下.包名(目录名称)/引用的包名

	_ "testmain/pkg"
)

//main函数是程序的入口,只写在main包里面
//开发者写代码的执行入口
func main() {

}

//初始化函数
//导入包的时候执行
//程序开始时会执行所有init函数
//一个文件里可以有多个init函数
//同一个包下的2个go文件中.都有init函数的话.是随机执行
//所以不要依赖init来执行程序的顺序
// init函数一般与"_"配合使用 来初始化一些驱动或者函数
func init() {
	fmt.Println("main.init.go")
}
