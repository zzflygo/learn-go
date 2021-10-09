package main

import "fmt"

/*
  基本数据类型:int, float, string, bool
  复合数据类型:array, slice, map, function, pointer, struct...


   arry:[len]type    arr:=[5]int
   slice:[]type      s1:=make([]int,0,5)    s1 = append(s1,数据1,数据2)
   map:map[key类型]value的类型   map1:=make(map[string]string)

   make创建的都是引用类型
*/
func main() {
	ap1 := make(map[int]string)
	map2 := make(map[int]map[string]string) //可以创建双重map key对应一个map

	fmt.Printf("%T \n %T \n", ap1, map2)

	s1 := []int{1, 3, 5, 7, 9}
	fmt.Println(s1)
	s1[0] = 66
	fmt.Println(s1)
	s1 = append(s1, 44, 56, 77, 6512)
	s1 = append(s1, 60)
	fmt.Println(s1)
	s1[9] = 35
	fmt.Println(s1)
	fmt.Println("~~~~~~~~~~~~~~~")
	s2 := s1

	//copy(s2, s1)
	//fmt.Println(s2)

	s2[3] = 111
	fmt.Println(s1)
	fmt.Println(s2)

}
