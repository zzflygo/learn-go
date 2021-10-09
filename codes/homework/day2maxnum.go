package main

import (
	"fmt"
)

func main() {

	/*var nums = []int{108, 107, 105, 109, 120, 103, 102, 99}
	sort.Ints(nums)
	fmt.Println(nums)
	fmt.Println(nums[len(nums)-1])
	*/
	var nums = []int{108, 107, 105, 109, 120, 103, 102, 99}
	for n := 0; n < len(nums); n++ {

		for i := 0; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}

		}
	}
	fmt.Println(nums)

}
