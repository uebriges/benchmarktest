package longestPalindromicSubstring

import (
	"fmt"
)

func longestPalindrome(s string) string {
	singleCharsCounter := map[string]int{}
	for _, v := range s {
		char := string(v)
		fmt.Printf("Value %s\n", char)
		singleCharsCounter[char] = singleCharsCounter[char] + 1
	}
	fmt.Println(singleCharsCounter)
	return s
}
