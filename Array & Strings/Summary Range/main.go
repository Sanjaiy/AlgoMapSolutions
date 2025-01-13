// https://leetcode.com/problems/summary-ranges/

package main

import (
	"fmt"
)

// func summaryRanges(nums []int) []string {
// 	temp := make([][]int, len(nums))
// 	var tempIndex int

// 	temp[tempIndex] = append(temp[tempIndex], nums[0])

//     for i, val := range nums {
// 		if i == 0 {
// 			continue
// 		}

// 		if i > 0 && val -1 == nums[i-1] {
// 			temp[tempIndex] = append(temp[tempIndex], val)
// 		} else {
// 			tempIndex++
// 			temp[tempIndex] = append(temp[tempIndex], val)
// 		}
// 	}

// 	op := make([]string, len(temp))
// 	var opIndex int

// 	for _, val := range temp {
// 		if len(val) == 0 {
// 			break
// 		}

// 		var opVal string

// 		if len(val) == 1 {
// 			opVal = fmt.Sprintf("%d", val[0])
// 		} else {
// 			opVal = fmt.Sprintf("%d->%d", val[0], val[len(val)-1])
// 		}

// 		op[opIndex] = opVal
// 		opIndex ++
// 	}

// 	return op
// }

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string {}
	}

	var op []string

	addRange := func (start, end int) {
		if start == end {
			op = append(op, fmt.Sprintf("%d", start))
		} else {
			op = append(op, fmt.Sprintf("%d->%d", start, end))
		}
	}

	start := nums[0] 
	end := start

	for i := 1 ; i < len(nums); i++ {
		if nums[i] != end + 1 {
			addRange(start, end)
			start = nums[i]
		}
		end = nums[i]
	}

	addRange(start, end)
	
	return op
}

func main() {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
	fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}