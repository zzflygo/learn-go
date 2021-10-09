package main

import "fmt"

func main() {

	for aa := 1; aa < 10; aa++ {

		for bb := 1; bb < 10; bb++ {
			fmt.Printf("%d * %d = %d \t ", aa, bb, aa*bb)
		}
	}
}
