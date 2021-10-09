package main

import "fmt"

func main() {
	// 1+2+3+....100
	sum := 0
	idx := 1
START: //标签
	if idx > 100 {
		goto END
	}
	sum += idx
	idx += 1
	goto START
END: //标签
	fmt.Println(sum)
}
