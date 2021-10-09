package main

import "fmt"

/* 可变参数
当一个函数的形参不确定具体的个数的时候,可以用可变参数.

在一个函数中 有多个形参时,可变参数有且只有一个.放在句子的最后面
(n...int)(n...float)

*/

func main() {
	getSum(7, 8, 9)
}

func getSum(n ...int) {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	fmt.Printf("和:%d", sum)
}
