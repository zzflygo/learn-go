package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int // json 包想要访问这个结构体 需要首字母大写
	Name string
}

type Class struct { //结构体标签Tag   必须是反引号``包裹的name:"name"
	Title    string    `json:"ss"`   //tag 可以把字段名改成别的包或者语言需要名字
	Students []Student `json:"love"` //格式为: ` json:"type" xml:"type2" db:"data"     ` 多个包用一个空格
}

func newStudent(id int, name string) Student {
	return Student{
		Id:   id,
		Name: name,
	}
}

func main() {
	//chi := make([]student, 0, 20)
	c1 := Class{
		Title:    "小葵花1班",
		Students: make([]Student, 0, 20),
		//Students: chi,
	}
	for i := 0; i < 10; i++ {
		tmpstu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpstu)
	}
	fmt.Printf("%s \n", c1)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~")
	//JSON序列化:GO语言中的数据 ->JSON格式的字符串
	d1, err := json.Marshal(c1)
	fmt.Println(err)
	if err != nil {
		fmt.Println("jrson.Marshal failed ,err", err)
		return
	}
	fmt.Printf("%T \n", d1)
	fmt.Printf("%s \n", d1)

	jsonStr := `{"ss":"小葵花1班","love":[{"Id":0,"Name":"stu00"},{"Id":1,"Name":"stu01"},{"Id":2,"Name":"stu02"},{"Id":3,"Name":"stu03"},{"Id":4,"Name":"stu04"},{"Id":5,"Name":"stu05"}]}`
	fmt.Printf("%T \n", jsonStr)
	var c2 Class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json umarshal err \n", err)
		return
	}
	fmt.Printf("%#v \n", c2)
}
