package main

import "fmt"

//类型断言 就是猜类型
//2种方式
// 1 .if 语句 v,ok:=a.(要猜的类型)
// 2 .switch语句 v:=a(type)-->type为关键字

func assign(a interface{}) {
	v, ok := a.(string)
	if !ok {
		fmt.Println("猜错了..")
	} else {
		fmt.Println("猜对了为string", v)
	}
}

func assign2(a interface{}) {
	switch t := a.(type) {
	case string:
		fmt.Printf("类型为%T->%v\n", t, t)
	case int:
		fmt.Printf("类型为%T->%v\n", t, t)
	case float64:
		fmt.Printf("类型为%T->%v\n", t, t)
	case bool:
		fmt.Printf("类型为%T->%v\n", t, t)
	default:
		fmt.Println("没有该类型")

	}
}

func main() {
	//assign("小西瓜")
	assign2(32)
	assign2("吸管刷")
	assign2(3.1415)
	assign2(false)

}
