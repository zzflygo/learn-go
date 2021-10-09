package main

import "fmt"

func main() {

	fmt.Println(sum(5))

}

func sum(n int) int {
	if n == 1 { //	闭包
		return 1
	}

	return sum(n-1) + n //累加
}
