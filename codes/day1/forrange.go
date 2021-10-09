package main

import "fmt"

func main() {
	s := "我爱中国"
	for i, v := range s {
		fmt.Printf("%v, %q\n", i, v)
	}
}
