package main

import "fmt"

//结构体狗
type dog struct {
	name string
	col  string
}

//构造结构体函数
func newDog(name, col string) *dog {
	return &dog{
		name: name,
		col:  col,
	}
}

//方法 是特定变量才能使用的函数
//接受者表示调用该方法的具体类型变量,多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪! \n", d.name)
}
func main() {
	d1 := newDog("大黄", "黄色")
	d1.wang()

}
