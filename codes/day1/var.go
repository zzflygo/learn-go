package main

import "fmt"
// 我是行注释
/*   
我是块注释
 */

// 全局变量
// var 定义变量
// var 标识符（变量名称） 类型（变量的类型）
// string 表示字符串
// 若未声明 则使用零值进行初始化

var name string = "kk"

func main(){
        // 局部变量
       // int 整数
        var age int = 31

        fmt.Println(name, age)
}