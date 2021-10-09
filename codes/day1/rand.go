package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Int() //random伪随机数 根据一定算法公式算出来的
	fmt.Println(num1)
	fmt.Printf("%T", num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10) //0~9之间 不包含10
		fmt.Println(num)
	}
	fmt.Println("####################")
	rand.Seed(2) //随机数的算法的种子要求int64
	for i := 0; i < 10; i++ {

		num2 := rand.Intn(10) //0~9之间 不包含10
		fmt.Println(num2)
	}
	fmt.Println("####################")

	t1 := time.Now()
	fmt.Println(t1)
	//时间戳:指定时间,距离1970年1月1日0点0分0秒,之间的时间差值:秒 ,纳秒
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)

	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2)

	//step1: 设置种子数, 可以设置成时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//stpe2: 调佣生成随机数的函数
		fmt.Println("-->", rand.Intn(100))
	}

	//[15,76]就是(0-61)+15  15-75的数
	//[3,48] 就是(0,45)+3   3-47的数
	num3 := rand.Intn(46) + 3
	fmt.Println(num3)
}
