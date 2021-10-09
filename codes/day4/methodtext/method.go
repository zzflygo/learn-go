package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

var (
	id   int
	name string
)
var class map[int]*student

func addStudent() *student {

	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入名字:")
	fmt.Scanln(&name)

	return &student{
		id:   id,
		name: name,
	}

}

func showStudent() {
	for k, v := range class {
		fmt.Printf("学号:%d 名字:%s \n", k, v.name)
	}
}

func main() {
	class = make(map[int]*student, 100)
	smr = studentMgr{
		allStudent: make(map[int]*student,100)
	}

	for {
		show()
	}

}

func modifyStudent() {
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入名字:")
	fmt.Scanln(&name)
	class[id] = &student{
		id:   id,
		name: name,
	}
}

func show() {
	var input int
	fmt.Println("    用户主菜单")
	fmt.Println("1. 添加学生信息")
	fmt.Println("2. 修改学生信息")
	fmt.Println("3. 学生信息展示")
	fmt.Println("4.   退出系统")
	fmt.Println("请输入数字进行选择")
	fmt.Scanln(&input)
	fmt.Println("您输入的选择为:", input)
	switch input {
	//增加学生信息
	case 1:
		class[id] = addStudent()
	//修改学生信息
	case 2:
		modifyStudent()
	//展示学生信息
	case 3:
		showStudent()

		//退出系统
	case 4:
		os.Exit(0)
	default:
		fmt.Println("滚!")
	}

}
