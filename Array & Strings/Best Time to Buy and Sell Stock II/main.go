package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var op int
	
	for i := 1; i < len(prices); i++ {
		op += max(0, prices[i] - prices[i-1])
	}

	return op
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}