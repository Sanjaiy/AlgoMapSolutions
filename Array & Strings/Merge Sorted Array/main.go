package main

import (
	"fmt"
	// "slices"
)

// Solution: 1
// func merge(nums1 []int, m int, nums2 []int, n int)  {
// 	var ptr int

//     for i := m ; i < m+n ; i++ {
// 		nums1[i] = nums2[ptr]
// 		ptr++
// 	}

// 	slices.Sort(nums1)
// }

// Solution: 2
func merge(nums1 []int, m int, nums2 []int, n int)  {
	ptr1 := m-1
	ptr2 := n-1
	ptr3 := m+n-1

    for ptr2 >= 0 {
		if ptr1 >=0 && nums1[ptr1] > nums2[ptr2] {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		}
		ptr3--
	}
}

func main() {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int {2,5,6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{1}
	m = 1
	nums2 = []int {}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{0}
	m = 0
	nums2 = []int {1}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}