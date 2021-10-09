package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		fmt.Printf("%p-->%d \n", &i, i)
		defer func(i int) {
			fmt.Printf("%p-->%d \n", &i, i) //在循环每次声明函数的时候 新创建了一个i申请一个新的内存地址
			fmt.Printf("defer %d \n", i)    //此时 i 是外部变量.defer把打印函数推到最后运行
			//.但是此时i已经i++到了3.所以最后打印出来的是3
		}(i)
		fmt.Printf("%p-->%d \n", &i, i)
	}

}
