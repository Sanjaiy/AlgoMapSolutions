package main

import (
	"fmt"
	"strings"
)

// Approach: 1
func lengthOfLastWord1(s string) int {
	str := strings.Split(strings.Trim(s," "), " ")

	return len(str[len(str) - 1])
}

// Approach: 2
func lengthOfLastWord2(s string) int {
    length := 0
    s = strings.TrimSpace(s)

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            break
        }
        length++
    }
    return length
}

func main() {
	fmt.Println(lengthOfLastWord1("Hello World"))
	fmt.Println(lengthOfLastWord1("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord2("luffy is still joyboy"))
}