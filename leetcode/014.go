package leetcode

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	res := ""
	for i := 1; i < len(s)+1; i++ {
		check := s[:i]
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], check) {
				return res
			}
		}
		res = check
	}
	return res
}
