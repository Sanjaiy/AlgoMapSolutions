package main

import (
	"fmt"
	"math"
)

func trap(height []int) int {
    var op int

    left, right := 1, len(height) - 2

    lMax, rMax := height[left-1], height[right+1]

    for left <= right {
        if rMax <= lMax {
            op += int(math.Max(0, float64(rMax)-float64(height[right])))

            rMax = int(math.Max(float64(rMax), float64(height[right])))

            right -= 1
        } else {
            op += int(math.Max(0, float64(lMax)-float64(height[left])))

            lMax = int(math.Max(float64(lMax), float64(height[left])))

            left += 1
        }
    }

    return op
}

func main() {
    fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
    fmt.Println(trap([]int{4,2,0,3,2,5}))
}