package main

import (
	"fmt"
)

func reverse(arr []int) {
	for i,j := 0, len(arr)-1; i < j; i,j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	} 
} 

func rotate(nums []int, k int)  {
	k = k % len(nums)
    reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main() {
	arr1 := []int{1,2,3,4,5,6,7}
	rotate(arr1, 3)
	fmt.Println(arr1)

	arr2 := []int{-1,-100,3,99}
	rotate(arr2, 2)
	fmt.Println(arr2)

	arr3 := []int{-1}
	rotate(arr3, 2)
	fmt.Println(arr3)

	arr4 := []int{-1, 2, 3}
	rotate(arr4, 3)
	fmt.Println(arr4)
}