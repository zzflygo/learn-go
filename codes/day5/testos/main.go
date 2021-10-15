package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] //os.Args[0]是存储程序的绝对地址
	var sum int
	uerse := func() {
		fmt.Println("usras 1 2 [3 4]")
	}

	if len(args) < 2 {
		uerse()
		return
	}

	for _, v := range args {
		if intValue, err := strconv.Atoi(v); err != nil {

			uerse()
			break
		} else {
			sum += intValue
		}
	}

	fmt.Println(sum)

}
