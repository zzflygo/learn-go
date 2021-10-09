package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	name    string
	tel     string
	address string
	id      string
	user    = []map[string]string{}
)

func main() {

	for {
		fmt.Println("----用户管理系统-----")
		fmt.Println("-----1.添加用户-----")
		fmt.Println("-----2.修改用户-----")
		fmt.Println("-----3.删除用户-----")
		fmt.Println("-----4.查找用户-----")
		fmt.Println("-----5.全部用户-----")
		fmt.Println("-----6.退出系统-----")

		chice := 0
		fmt.Scan(&chice)
		switch {
		case chice == 1:
			add()
		case chice == 2:
			modify()
		case chice == 3:
			del()
		case chice == 4:
			query()
		case chice == 5:
			if len(user) == 0 {
				fmt.Println("无数据")
			} else {
				fmt.Println(user)
			}
		case chice == 6:
			return
		default:
			fmt.Println("请输入 1-5 ")

		}

	}
}
func add() {

	userInfo := map[string]string{"id": " ", "name": " ", "tel": " ", "address": " "}

	fmt.Println("请输入名字:")
	fmt.Scan(&name)
	fmt.Println("请输入电话:")
	fmt.Scan(&tel)
	fmt.Println("请输入地址:")
	fmt.Scan(&address)

	if len(user) == 0 {
		id = "1"
	} else {
		id = strconv.Itoa(len(user) + 1)
	}
	userInfo["id"] = id
	userInfo["name"] = name
	userInfo["tel"] = tel
	userInfo["address"] = address
	user = append(user, userInfo)

}

func del() {
	var n int
	var is string
	fmt.Println("请输入想要删除的id: ")
	fmt.Scan(&n)
	fmt.Println("请确认删除的信息:%q ,删除请输入: y \n", user[n-1])
	fmt.Scan(&is)
	if is != "y" {
		fmt.Println("退出系统")
		return
	}

	s := func() {
		if n == 1 {
			user = user[1:]
		} else if n == len(user) {
			user = user[:len(user)-1]
		} else if n > 1 && n < len(user) {
			user = append(user[:n-1], user[n:0]...)
		} else if n > len(user) {
			fmt.Println("不存在此用户")
		}
	}
	s()
	fmt.Println(user)
}
func modify() {
	var n int
	var is string
	fmt.Println("请输入想要修改的id: ")
	fmt.Scan(&n)
	fmt.Printf("请确认修改前的信息:%q ,确认修改请输入: y \n", user[n-1])
	fmt.Scan(&is)
	if is != "y" {
		fmt.Println("退出系统")
		return
	}
	fmt.Println("请输入修改的名字:")
	fmt.Scan(&name)
	fmt.Println("请输入修改的电话:")
	fmt.Scan(&tel)
	fmt.Println("请输入修改的地址:")
	fmt.Scan(&address)

	user[n-1]["name"] = name
	user[n-1]["tel"] = tel
	user[n-1]["adress"] = address
	fmt.Println("修改后的内容为: ", user[n-1])
}

func query() {
	var que string
	fmt.Println("请输入查询内容: ")
	fmt.Scan(&que)
	for i, n := range user {

		for _, p := range n {
			if strings.Contains(p, que) {

				fmt.Printf("查找结果为 :\n ID: %s  名字:%s  地址: %s  电话:%s \n", user[i]["id"], user[i]["name"], user[i]["address"], user[i]["tel"])

			} else {
				fmt.Println("无此信息")
			}
		}
	}

}
