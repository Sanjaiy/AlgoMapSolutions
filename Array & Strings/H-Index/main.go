package main

import "fmt"

func hIndex(citations []int) int {
    var op int
	n := len(citations)
	counts := make([]int, n+1)

	for _,citation := range citations{
		if citation > n {
			counts[n]++
		} else {
			counts[citation]++
		}
	}

	for i:=n; i >=0;i--{
		op += counts[i]

		if op >= i {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(hIndex([]int{3,0,6,1,5}))
	fmt.Println(hIndex([]int{1,3,1}))
}