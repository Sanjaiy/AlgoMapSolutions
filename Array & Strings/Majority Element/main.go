package main

import "fmt"

func majorityElement(nums []int) int {
    var op int
	size := len(nums)/2
	counts := make(map[int]int)
	
	for _, val := range nums {
		counts[val]++
	}

	for k,v := range counts {
		if v > size {
			op = k
			size = v		
		}
	}

	return op
}

func main() {
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}