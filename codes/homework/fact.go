package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	fmt.Println(mul(a))

}

func mul(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n < 0 {
		return -1
	}
	return mul(n-1) * n
}
