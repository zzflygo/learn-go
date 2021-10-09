package main

import "fmt"

type animal interface {
	move()
	eat(s string)
}

type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println(c.name, "走猫步..")
}
func (c cat) eat(x string) {
	fmt.Println(c.name, "在吃", x)
}

type flsh struct {
	name string
	feet int
}

func (f flsh) move() {
	fmt.Println(f.name, "在游泳")
}
func (f flsh) eat(q string) {
	fmt.Println(f.name, "在吃", q)
}
func main() {
	var a1 animal
	c1 := cat{
		name: "多多",
		feet: 4,
	}
	a1 = c1
	a1.move()
	a1.eat("鱼")

	f1 := flsh{
		name: "晶晶",
		feet: 0,
	}
	a1 = f1
	a1.move()
	a1.eat("蚯蚓")

}
