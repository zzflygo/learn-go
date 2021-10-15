package mathadd

import (
	"fmt"
	"strconv"
)

func Add(k []string) int {

	n := 0
	m, err := strconv.Atoi(k)
	if err != nil {
		fmt.Println("需要输入2个以上整数")
		return -1
	}
	if len(m) < 2 {
		fmt.Println("至少输入2个整数~~")
		return -1
	}
	for _, v := range m {
		n += v
	}
	return n

}
