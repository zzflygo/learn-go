package main

import "fmt"

/* map : 映射, 是一种撞门用来储存键值对的集合.属于引用类型.

存储特点:
	A:存储是无序的键值对
	B:键值不能重复, 并且和value值 一一对应的.
	    	map中的key不能重复,如果重复了.那么行的value会覆盖原来的

语法结构:
	1.创建map
		var map1 map[key类型]value类型.此时创建的是nil map,无法直接使用.

		var map2 = make(map[key类型]value类型)

		var map3 = make[key类型]value{key1:value1,key2:value2,key3:value3.....}

	2.添加/修改
		map[key] = value
		如果key不存在则是添加
		如果key存在则是修改
	3.获取
		map[key]--->value
		value,pk:=map[key]
			根据key获取对应的value
				如果key存在,value就是对应的数据,ok为true
				如果key不存在,value就是值类型的默认值,ok为false
	4.删除数据:
		delete(map,key)
			如果key存在, 就可以直接删除
			如果key不存在, 删除失败, 无影响
	5.长度
		len(map) 在获取有一个key

*/

func main() {
	//创建map
	var map1 map[int]string // 没有初始化,nil 不能直接使用
	var map2 = make(map[int]string)
	var map3 = map[string]int{"go": 98, "java": 88, "python": 79, "html": 95}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil) // nil 等于 空. map为空的时候 不能直接使用
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)
	// nil map
	if map1 == nil {
		map1 = make(map[int]string)
	}
	fmt.Println(map1 == nil) //为nil的map,用map初始化以后就不为空了

	//2.添加/修改
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "memeda"
	map1[4] = "王二狗"
	map1[5] = "ruby"
	map1[6] = "思密达"
	map1[7] = ""

	fmt.Println(map1)
	fmt.Println(map1[4])
	fmt.Println(map1[40]) //key不存在时,得到value默认格式的零值.

	s1, ok := map1[6] // 把key对应的值,赋值给了新的变量
	if ok {
		fmt.Println("对应的数值是: ", s1)
	} else {
		fmt.Println("操作key不存在,得到零值 ", s1)
	}
	fmt.Println(s1)
	fmt.Printf("%T \n", s1)
	//更改
	map1[3] = "小giegie"
	fmt.Println(map1[3])
	map1[8] = "抖音"
	fmt.Println(map1)
	delete(map1, 4)
	fmt.Println(map1)
	fmt.Println(len(map1))

}
