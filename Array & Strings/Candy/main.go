package main

import "fmt"

func candy(ratings []int) int {
    if len(ratings) == 1 {
		return 1
	}

	var op int
	candies := make([]int, len(ratings))
	for i := range candies {
        candies[i] = 1
    }

	for i:=1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i:= len(ratings)-2; i >= 0; i--{
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}

		op += candies[i]
	}
	op += candies[len(ratings)-1] 

	return op
}

func main() {
	fmt.Println(candy([]int{1,0,2}))
	fmt.Println(candy([]int{1,2,2}))
}