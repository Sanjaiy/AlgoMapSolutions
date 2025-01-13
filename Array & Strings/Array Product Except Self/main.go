// https://leetcode.com/problems/product-of-array-except-self/

package main

import "fmt"

// func productExceptSelf(nums []int) []int {
// 	if len(nums) == 0 {
// 		return []int {}
// 	}

// 	product := func (arr []int) int {
// 		mul := 1
// 		for _, val := range arr {
// 			mul *= val
// 		}

// 		return mul
// 	}

// 	var op []int
	
// 	for i := range nums {
// 		temp := append(nums[i+1:], nums[:i]...)
// 		op = append(op, product(temp))
// 	}

// 	return op
// }


func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int {}
	}
	
	op := make([]int, len(nums))
	leftMul := 1
	rightMul := 1

	for i := 0; i < len(nums); i++ {
		op[i] = leftMul
		leftMul *= nums[i]
	}

	for i := len(nums)-1; i >=0; i-- {
		op[i] *= rightMul
		rightMul *= nums[i]
	}

	return op
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3}))
}