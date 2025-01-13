// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	i := 0
	j := len(prices) - 1
	op := 0

	for idx, val := range prices {
		if prices[i] > val {
			i = idx
		}

		if idx > i && val > prices[j] {
			j = idx
		}

		if j > i {
			profit := prices[j] - prices[i]
			if profit > op {
				op = profit
			}
		}
	}

	return op
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{7,10}))
	fmt.Println(maxProfit([]int{10,6,5,1,3}))
	fmt.Println(maxProfit([]int{2,4,1}))
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
}