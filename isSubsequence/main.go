// https://leetcode.com/problems/is-subsequence/

package main

import (
	"fmt"
	"unicode/utf8"
)


func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	chars := []rune(s)
	var j int

	if len(chars) > utf8.RuneCountInString(t) {
		return false
	}

	for _, val := range t {
		if val == chars[j] {
			j += 1

			if j == len(chars) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("ace", "abcde"))
	fmt.Println(isSubsequence("aec", "abcde"))
}