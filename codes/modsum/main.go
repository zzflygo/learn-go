package main

import (
	"fmt"
	"modsum/mod1/ssr"
	"modsum/mod2"
	"modsum/mod3"
	"testmod"
)

func main() {
	//mod2.M2()
	mod2.M2()
	mod3.M3()
	ssr.Ss1()
	s := testmod.Mul(2, 4)
	fmt.Println(s)

}
