package main

import (
	"fmt"
)

func main() {

	tower("A", "B", "C", 3)
}

//把n 盘子 由a 经过b 传到 c
func tower(a, b, c string, n int) {
	if n == 1 {
		fmt.Println(a, "-->", c)
		return
	}
	//第一步 把n-1个盘子 由a 经过c 传到b
	tower(a, c, b, n-1)
	//第二步,把第n个盘子直接放到c
	fmt.Println(a, "-->", c)
	//第三部,把n-1个盘子由b 经过 a 传到c
	tower(b, a, c, n-1)

}
