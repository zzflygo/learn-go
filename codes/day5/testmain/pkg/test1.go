package pkg

import (
	"fmt"
	_ "testmain/pkg/ssr"
)

func init() {
	fmt.Println("test1.init.go")
}
func init() {
	fmt.Println("test2.init.go")
}
