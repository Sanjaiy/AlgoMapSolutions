// https://leetcode.com/problems/longest-common-prefix/

package main

import (
	"fmt"
	// "strings"
)

// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
//         return ""
//     }

//     minString := strs[0]

//     for _, str := range strs {
//         if len(minString) > len(str) {
// 			minString = str
// 		}
//     }

// 	if len(minString) == 0 {
// 		return ""
// 	}

// 	var op string

// 	for i := 1; i <= len(minString) ; i++ {
// 		z := minString[:i]
// 		for _, str := range strs {
// 			if str == minString {
// 				continue
// 			}
// 			if !strings.HasPrefix(str, z) {
// 				return op
// 			}
// 		}
// 		op = z
// 	}

// 	return op
// }


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    reference := strs[0]
    
    for i := 0; i < len(reference); i++ {
        for _, str := range strs[1:] {
            if i >= len(str) || str[i] != reference[i] {
                return reference[:i]
            }
        }
    }
    
    return reference
}


func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"r","racecar","rr"}))
}