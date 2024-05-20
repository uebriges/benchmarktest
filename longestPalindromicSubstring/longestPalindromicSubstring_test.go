package longestPalindromicSubstring

import (
	"fmt"
	"testing"
)

func Test_Lpst(t *testing.T) {
	res := longestPalindrome("babad")
	fmt.Println(res)
	// assert.Equal(t, res, "bab")

	// res = longestPalindrome("cbbd")
	// assert.Equal(t, res, "bb")

}
