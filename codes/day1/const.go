package main

import "fmt"

const statusNew int = 1
const statusDeleted int = 2

func main() {
	const (
		Monday  = 1
		Tuesday = 2
	)
	fmt.Println(statusNew, statusDeleted)
	fmt.Println(Monday, Tuesday)
	fmt.Printf("%T %T %T %T", statusNew, statusDeleted, Monday, Tuesday)
}
