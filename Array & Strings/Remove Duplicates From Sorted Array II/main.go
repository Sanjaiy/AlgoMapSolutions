package main

import "fmt"

func removeDuplicates(nums []int) int {
	op := 1
	lastVal := nums[0]
	lastValOccurence := 1

	for i, val := range nums {
		if i == 0 {
			continue
		}
		if val != lastVal {
			lastVal = val
			lastValOccurence = 1
			nums[op] = val
			op++
		} else {
			lastValOccurence++
			if lastValOccurence <= 2 {
				nums[op] = val
				op++
			} 
		}
	}

	return op
}

func main() {
	arr1 := []int {1,1,1,2,2,3}
	fmt.Println(removeDuplicates(arr1))
	fmt.Println(arr1)

	arr2 := []int {0,0,1,1,1,1,2,3,3}
	fmt.Println(removeDuplicates(arr2))
	fmt.Println(arr2)
}