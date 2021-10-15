package writebybufio

import (
	"bufio"
	"fmt"
	"os"
)

func Add(n string) {
	file1, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file failed err:", err)
		return
	}
	defer file1.Close()
	reader := bufio.NewWriter(file1)
	//写进缓存
	_, err = reader.WriteString(n)
	if err != nil {
		fmt.Println("open file failed err:", err)
		return
	}
	//写进内容
	reader.Flush()
}
