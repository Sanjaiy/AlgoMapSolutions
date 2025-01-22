package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	var op int

    for i:=0; i < len(nums); i++ {
		if op < i {
			return false
		}

		op = max(op, nums[i]+i)
	}

	return true
}

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{2,0}))
}