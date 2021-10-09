package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //设置一个种子 rand.Seed()  time.Now().Unix() 设置当前时间为种子
	f := say()
	f()

	f1 := strings.FieldsFunc("assdeesadaxxzcadsradd", aFields)
	fmt.Printf("%q \n", f1)

	e := strings.FieldsFunc("AscnaxksAxaxckaedfwadx", func(ch rune) bool {
		return ch == 'a'
	})
	fmt.Println(e)
}

func sayhi() {
	fmt.Println("hi")
}

func sayhello() {
	fmt.Println("hello")
}

func say() func() {
	if rand.Int()%2 == 0 { //设置一个随机整数 rand.Int()
		return sayhi
	}
	return sayhello
}

func aFields(split rune) bool {
	if split == 'a' { //可以直接写 return aplit == 'a' 自用一行
		return true
	}
	return false

}
