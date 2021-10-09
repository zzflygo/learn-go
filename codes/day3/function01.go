package main

import "fmt"

/*一, 概念:
	具有特殊功能的代码,可以被多次调用执行.

  二, 意义:
	1,可以避免重复的代码
	2,增强程序的扩展性
  三,使用
	step1,函数的定义,也叫声明 func 函数名(){}
	step2,函数的调用,就是执行函数中的代码
  四,函数的结构 func name(形体值,形体值)(返回值1,返回值2){
	  函数体,代码段
  }
  	没有形体值时可以省略不写,当返回值,没有或 只有1个时也可以不写
  五,注意
  		A,函数必须先定义,在调用,如果不定义:underfind:getSum
		  定义了函数,没有调用,那么函数就失去意义
		 B,函数名不能冲突
		 C,main()是特殊的函数,由系统自动调用.其他函数都是通过名字来调用.

*/

func main() { //程序入口,是一个特殊的函数
	getSum()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~")
	getSum()
	fmt.Println("aaaaaaaaaaa")
	getSum()
	fmt.Println("bbbbbbbbbbb")

}

func getSum() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
