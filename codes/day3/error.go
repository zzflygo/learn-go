package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	//需要给调用者返回错误信息
	//通过函数最后一个返回值返回错误信息
	//调用者需要对错误进行检查,决定自己如何操作
	//建议程序员去处理错误
	v, err := strconv.Atoi("123asd")
	if err == nil {
		fmt.Println("value是", v, "没有错误")
	} else {
		fmt.Println("输入错误")
	}
	//错误
	//运行时错误
	//			可恢复的错误 ==> 重试/忽略
	//			不可恢复的错误 ==> 程序退出

	if s, err := div(1, 0); err == nil {
		fmt.Println(s)
	} else {
		fmt.Println("value is error", err)
	}

}

// error 是一个接口. 通过fmt.Errorf 返回
func div(left, right int) (int, error) {
	if right == 0 {
		return 0, fmt.Errorf("right num is zero") //用fmt包来返回,效果一样
	}
	return left / right, nil
}

func div2(left, right int) (int, error) {
	if right == 0 {
		return 0, errors.New("right num is zero") //用errors包来返回,效果一样
	}
	return left / right, nil
}
