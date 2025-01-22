package main

import (
	"fmt"
)

// func jump(nums []int) int {
//     var op int

// 	for i,val := range nums {
// 		maxVal := max(nums[i:val+i])
// 		op = op + (maxVal - i)
// 	}

// 	return op
// }

// func max(nums []int) int {
// 	maxVal := nums[0]
// 	for _, val := range nums {
// 		if val > maxVal {
// 			maxVal = val
// 		}
// 	}
// 	return maxVal
// }

func jump(nums []int) int {
    var op, maxVal, curr int
	
	for i:=0; i < len(nums)-1; i++ {
		maxVal = max(maxVal, i+nums[i])

		if curr == i {
			op++
			curr = maxVal
		}
	}

	return op
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
}