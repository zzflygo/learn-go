package main

import "fmt"

type cat struct {
	name string
	feet int
}
type animal interface {
	runer
	eater
}
type runer interface {
	run()
}
type eater interface {
	eat(string)
}

func (c *cat) eat(s string) {
	fmt.Println(c.name, "吃", s)
}
func (c *cat) run() {
	fmt.Println(c.name, "pao~~")
}

type chicken struct {
	name string
	feet int
}

func (c chicken) eat(s string) {
	fmt.Println(c.name, "吃", s)
}
func (c chicken) run() {
	fmt.Println(c.name, "run~~")
}
func main() {
	var c1 = cat{
		name: "大黄",
		feet: 4,
	}
	var j1 = chicken{
		name: "唧唧",
		feet: 2,
	}
	var ss animal
	var ssr animal
	ss = &c1
	ssr = &j1
	ss.run()
	fmt.Printf("%T\n", ss)
	ssr.eat("米")
	fmt.Printf("%T\n", ssr)
}
