package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := i / 10 % 10
		z := i % 10
		if x*x*x+y*y*y+z*z*z == i {
			fmt.Print(i)

		}
	}
}
	for i:=2;i<=100;i++{
		
	}
