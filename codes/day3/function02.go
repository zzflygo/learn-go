package main

import "fmt"

/*
函数的参数:
	1,形式参数.定义参数时,用于接受外部传入的数据.
		一般在为 函数中不确定数值的变量,需要外部输入

	2,实际参数,为调用函数时,实际传给形参的数据.
	函数名必须匹配,实参和形参必须一一对应:顺序,个数,类型.const
*/

func main() {
	getSum(10, 0)

	getSum(20, 0)

	getSum(100, 0)

	func1(66, 98, "小橘子")
	func2(66.335, 9.8145, "小橘子")

	s1, s2 := rec(5, 6)
	fmt.Println("面积:", s1, "周长", s2)

}

func getSum(n int, sum int) { //int类型一致可以写在后边 func getSum(n ,m int)

	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1到%d的和:%d \n", n, sum)
}

func func1(a, b int, c string) {
	fmt.Printf("a:%d b:%d c:%s \n", a, b, c)
}
func func2(a, b float64, c string) {
	fmt.Printf("a:%.2f b:%.2f c:%s \n", a, b, c)
}

func rec(l, w float64) (float64, float64) {
	s := l * w
	d := (l + w) * 2
	return s, d
}
