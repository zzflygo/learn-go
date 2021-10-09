package main

import "fmt"

func main() {
	names := []string{"在", "输", "想", "你", "的", "液", "!"}
	fmt.Println(names)
	print(sy, names...)
	for i, v := range names {
		names[i] = sy(v)
	}
	fmt.Println(names)

}

func print(as func(string) string, args ...string) {
	for i, v := range args {
		fmt.Println(i, as(v))
	}
}

func sy(txt string) string {
	return "$" + txt + "$"
}
