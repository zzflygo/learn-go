package main

import "fmt"

func main() {
	/* map和slice的结合使用:
	1.创建map用于存储人的信息
		name, age, sex, address

	2.每个map存储一个人的信息

	3.将这些map存入到slice中

	4.打印遍历输出
	*/
	//第一个人
	map1 := make(map[string]string)
	map1["name"] = "王二狗"
	map1["age"] = "30"
	map1["sex"] = "男性"
	map1["address"] = "北京市XX路XX号"
	fmt.Println(map1)

	//第二个人
	map2 := map[string]string{"name": "李小花", "age": "20", "sex": "女", "address": "上海市.."}
	fmt.Println(map2)

	//第三个人

	var map3 = map[string]string{"name": "ruby", "age": "30", "sex": "女性", "address": "杭州"}
	fmt.Println(map3)

	//将map存到slice中
	s1 := make([]map[string]string, 0, 3) //切片只能用 append来添加元素
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	fmt.Println(s1)

	//遍历切片
	for i, val := range s1 {
		//val: map1,map2,map3
		fmt.Printf("第%d 个人的信息是: \n", i+1)
		fmt.Printf("\t 姓名 : %s \n", val["name"])
		fmt.Printf("\t 年龄 : %s \n", val["age"])
		fmt.Printf("%T", val)
	}

}
