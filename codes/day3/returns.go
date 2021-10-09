package main

import (
	"fmt"
)

func main() {
	fmt.Println(op(4, 6))
	a, b, c, d := op(4, 6)
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)
	nums := []int{122, 99, 66, 552, 12, 3, 33}
	ar := so(4, nums...)
	fmt.Println(nums)
	fmt.Println(ar)

}

func op(a, b int) (int, int, int, int) {

	return a + b, a - b, a * b, a / b

}

func so(b int, a ...int) []int {
	fmt.Printf("a的类型:%T a的值:%v \n", a, a)
	for n := 1; n <= b; n++ {
		for i := 0; i < len(a)-n; i++ {
			if a[i] >= a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			} else if a[i] == a[i+1] {
				continue
			}
		}
	}
	return a
}

func so1(a int, b func()) {
	for i := a - 1; a >= 0; i-- {
		b()
	}

}

func fbo(n int) int {

	for i := n; i >= 0; i-- {
		if fbo(1) == 1 || fbo(0) == 1 {

			return 1
		}
		return fbo(n-1) + fbo(n-2)
	}
}
