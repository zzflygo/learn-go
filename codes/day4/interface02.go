package main

import "fmt"

type falali struct {
	logo string
}

func (f falali) run() {
	fmt.Println(f.logo, "开到了90")
}

type baoshijie struct {
	logo string
}

func (b baoshijie) run() {
	fmt.Println(b.logo, "开到了130")
}

type kaiche interface {
	run()
}

func db(k kaiche) {
	k.run()
}

func main() {
	var f1 = falali{"f91"}
	var b1 = baoshijie{"911"}
	f1.run()
	b1.run()

	db(f1)
	db(b1)

}
