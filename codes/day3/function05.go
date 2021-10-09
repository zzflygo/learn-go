package main

import (
	"fmt"
)

/* 函数的返回值:
一个函数的执行结果, 返回给函数的调用处,执行结果就叫做函数的返回值

return语句:
	一个函数的定义上有返回值,那么函数中必须使用return语句,将结果返回给调用处
*/

func main() {
	a1 := getSum() //把返回的值赋予给a1 也可以直接使用
	//相当于 a1:=sum
	fmt.Println(getSum())
	fmt.Printf("%T \n", getSum())
	fmt.Println(a1)

	fmt.Println(getSum2())
	fmt.Printf("%T \n", getSum2())

	s1, c1 := rec(7, 9)
	fmt.Printf("面积:%.2f 周长:%.2f \n", s1, c1)

	s2, c2 := rec(4, 8)
	fmt.Printf("面积:%.2f 周长:%.2f \n", s2, c2)

}

func getSum() int { //只有一个返回值 名字可以写可以不写
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum //没写返回值的名字.则需要在return语句出填写返回的变量名

}
func getSum2() (vs int) { //vs 已经在返回函数上定义了.
	//不用在函数体中再定义一次 vs:=0
	for i := 1; i <= 100; i++ {
		vs += i
	}
	return //return后面是谁就把谁赋值给调用处
	//return 1 等于 return vs = 1 把1赋值给return再返回给调用处
}

func rec(len, wid float64) (float64, float64) {
	s := len * wid
	c := (len + wid) * 2
	return s, c
}

func rec2(len, wid float64) (s float64, c float64) {
	s = len * wid
	c = (len + wid) * 2
	return
}
