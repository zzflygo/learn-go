package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
ST:
	rand.Seed(time.Now().Unix()) //设置随机数种子
	var text int
	aa := rand.Intn(100)
	var yy string
	for i := 10; i >= 0; i-- {
		if i > 0 {
			fmt.Println("请输入一个数组:  还剩余", i-1, "次机会")
			fmt.Scan(&text)
			if aa < text {
				fmt.Println("输入的数字大了 ")
			} else if aa > text {
				fmt.Println("输入的数字小了 ")
			} else if aa == text {
				fmt.Println("正确了 ")
				goto AG
			}
		} else {
			fmt.Println("太笨了 ")
			goto AG
		}
	}
AG:
	fmt.Println("再来一次吗请输入 y ")
	fmt.Scan(&yy)
	if yy == "y" {
		fmt.Println("====================== ")
		goto ST
	} else {
		fmt.Println(" ")
	}
}
