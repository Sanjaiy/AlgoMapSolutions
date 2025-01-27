package main

import (
	"fmt"
	"math"
)

var intNumerals = map[int]string{
    1: "I",
    5: "V",
    10: "X",
    50: "L",
    100: "C",
    500: "D",
    1000: "M",
}

func firstDigit(num int) int {
	power := int(math.Log10(float64(num)))
	divisor := int(math.Pow(10, float64(power)))
	
	return num / divisor
}

func getMaxValue(num int) int {
	if num >= 1000 {
		return 1000
	} else if num >= 500 {
		return 500
	} else if num >= 100 {
		return 100
	} else if num >= 50 {
		return 50
	} else if num >= 10 {
		return 10
	} else if num >= 5 {
		return 5
	} else {
		return 1
	}
}

func getFirstMaxValue(num int) (int, int) {
	if num <= 1 {
		return 0, 1
	} else if num > 1 && num <= 5 {
		return 1, 5
	} else if num > 5 && num <= 10 {
		return 1, 10
	} else if num > 10 && num <= 50 {
		return 10, 50
	} else if num > 50 && num <= 100 {
		return 10, 100
	} else if num > 100 && num <= 500 {
		return 100, 500
	} else {
		return 100, 1000
	}
}

func intToRoman(num int) string {
    var op string

	for num > 0 {
		i := firstDigit(num)

		if num >= 1000 {
			op += intNumerals[1000]
			num -= 1000
		} else if i == 4 || i == 9 {
			j, k := getFirstMaxValue(num)
			op += intNumerals[j]
			op += intNumerals[k]
			num = num - (k-j)
		} else {
			maxVal := getMaxValue(num)
			op += intNumerals[maxVal]
			num -= maxVal
		}
	}

	return op
}

func main() {
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}