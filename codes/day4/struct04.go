package main

import "fmt"

func main() {
	d1 := Book{
		bookName:  "西游记",
		bookPrice: 46.5,
	}
	s1 := Student{name: "画画", age: 18, book: d1}
	fmt.Println(s1.name, s1.age)
	fmt.Println("\t", s1.book.bookName, s1.book.bookPrice)

	s2 := Student{name: "小灰灰", age: 32, book: Book{bookName: "十万个为什么", bookPrice: 25.3}}
	fmt.Println(s2.name, s2.age)
	fmt.Println("\t", s2.book.bookName, s2.book.bookPrice)

	s3 := Student{
		name: "张小峰",
		age:  35,
		book: Book{
			bookName:  "海底两万里",
			bookPrice: 128.63,
		},
	}
	fmt.Printf("name:%s,age:%d,book:%s,price:%.2f \n", s3.name, s3.age, s3.book.bookName, s3.book.bookPrice)

	d2 := Book{
		bookName:  "小猪佩奇",
		bookPrice: 33.65,
	}

	s4 := Student2{"小月", 20, &d2}
	fmt.Println(s4.name, s4.age)
	fmt.Println("\t", s4.book.bookName, s4.book.bookPrice)
	s4.book.bookName = "汪汪队"
	fmt.Println(d2.bookName)
}

type Book struct {
	bookName  string
	bookPrice float64
}

type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book // 结构体指针使用是不会发生值传递.就不会开辟内存空间,减少内存的使用
}
