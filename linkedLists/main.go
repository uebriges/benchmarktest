package main

import (
	"fmt"
	"slices"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	s4 := " "

	l1 := lengthOfLongestSubstring(s1)
	l2 := lengthOfLongestSubstring(s2)
	l3 := lengthOfLongestSubstring(s3)
	l4 := lengthOfLongestSubstring(s4)

	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)

	intslice := []int{1, 2, 3}
	// fmt.Println("length of space: ", len(space))
	// fmt.Println(strings.Contains(space, " "))
	slices.Contains[int](intslice, 3)
}

func lengthOfLongestSubstring(s string) int {
	rS := ""
	sL := len(s)
	top := 0
	if sL == 1 {
		return 1
	} else if sL == 0 {
		return 0
	}
	for i, _ := range s {
		cSS := string(s[i:])
		if len(cSS) < top {
			break
		}
		for _, cSSL := range cSS {
			if !strings.Contains(rS, string(cSSL)) {
				rS += string(cSSL)
			} else {
				if top < len(rS) {
					top = len(rS)
				}
				rS = ""
				break
			}
		}
	}
	return top
}

// var sizes []int
// sizes = append(sizes, len(rS))
// sort.Ints(sizes)
// return sizes[len(sizes)-1]
