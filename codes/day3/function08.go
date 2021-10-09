package main

import (
	"fmt"
)

func main() {
	a, b := swap(6, 9)

	fmt.Println(a, b)

}

func swap(n1, n2 int) (a1, a2 int) {
	n1, n2 = n2, n1 //交换
	return n1, n2
}
