// https://leetcode.com/problems/roman-to-integer/

package main

import (
	"fmt"
)

func roman(s string) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var op int
	var prev rune

	for i, roman := range s {
		op += romanNumerals[roman]

		if i > 0 {
			switch roman {
				case 'V', 'X' :
					if prev == 'I' {
						op -= 2
					}
				case 'L', 'C': 
					if prev == 'X' {
						op -= 20
					}
				case 'D', 'M':
					if prev == 'C' {
						op -= 200
					}
			}
		} 	
		prev = roman
    }

	fmt.Println(op)
}


func main() {
	roman("III")
	roman("IV")
	roman("IX")
	roman("VII")
	roman("XL")
	roman("XC")
	roman("CD")
	roman("CM")
	roman("CII")
}