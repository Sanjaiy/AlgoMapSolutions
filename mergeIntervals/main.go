// https://leetcode.com/problems/merge-intervals/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int {}
	}

	var op [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	op = append(op, intervals[0])
	var opIndex int

	for i, val := range intervals {
		if i == 0 {
			continue
		}

		if op[opIndex][1] >= val[0] {
			if op[opIndex][1] < val[1] {
				op[opIndex][1] = val[1]
			}
			if op[opIndex][0] > val[0] {
				op[opIndex][0] = val[0]
			}
		} else {
			op = append(op, val)
			opIndex ++
		}
	}

    return op
}

func main() {
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{1,4},{4,5}}))
	fmt.Println(merge([][]int{{1,4},{0,4}}))
	fmt.Println(merge([][]int{{1,4},{0,1}}))
	fmt.Println(merge([][]int{{1,4},{0,0}}))
}