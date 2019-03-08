package leetcode

import (
	"strings"
)

func LongestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	n := 0
	i := strings.Index(s)
}
