package main

import "fmt"

func main() {
	// 1 + 2 + 3 +......+100
	// for x; y ;z{    x
	//			N
	//}
	s := 0
	for i := 1; i <= 100; i++ {
		s += i
	}
	fmt.Println(s)

	// break 终端循环
	// contine 是跳过本次 继续下一次
}
