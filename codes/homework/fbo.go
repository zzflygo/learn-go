package main

import "fmt"

func main() {
	a := fbo(5)
	fmt.Println(a)
}

func fbo(n int) int {

	if n == 1 || n == 2 {

		return 1
	}
	return fbo(n-1) + fbo(n-2)
}
