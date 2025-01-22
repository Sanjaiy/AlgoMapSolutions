package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	
	i,j := n-1, n-1
	total, visited := 0, 0

	for visited < n {
		total += gas[j] - cost[j]
		visited += 1

		j = (j + 1) % n

		for total < 0 && visited < n {
			i = (i - 1 + n) % n

			total += gas[i] - cost[i]
			visited += 1
		}
	}

	if total < 0 {
		return -1
	} else {
		return i
	}
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))
}