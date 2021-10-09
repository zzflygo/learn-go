package main

import "fmt"

/*  1.模拟继承性  is - a
	父结构体
		子结构体

2.模拟聚合关系 has -a
*/
type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	school string
}

func main() {

	p1 := Person{name: "老张", age: 60, sex: "男"}
	fmt.Println(p1)

	s1 := Student{Person: Person{"小花", 18, "女"}, school: "昆工"}
	fmt.Println(s1.name, s1.Person.age, s1.sex, s1.school)

	s2 := Student{Person{"小灰灰", 30, "男"}, "广西大学"}
	fmt.Println(s2.name, s2.Person.age, s2.sex, s2.school)
	/* s3.Person.name --->s3.name
	Student结构体中使用了Person结构体作为匿名字段了
	那么Person中的字段 对于Student来说 就是 提升字段
	Student对象可以直接访问Person的字段,s3.Person.name 可简写为s3.name

	如果Student中使用了Person作为字段,并且命名的话,
	则不能简写

	*/
	var s3 Student
	s3.name = "老亮" //
	s3.age = 35
	s3.sex = "男"
	s3.school = "云艺"
	fmt.Println(s3)

}
