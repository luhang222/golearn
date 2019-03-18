package leetcode

import "strconv"

/**
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/
func CountAndSay(n int) string {
	res := "1"
	for i := 0; i < n-1; i++ {
		res = translate(res)
	}
	return res
}

func translate(s string) string {
	res := ""
	c := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			res += strconv.Itoa(count) + string(c)
			c = s[i]
			count = 1
		} else {
			count++
		}
		if i == len(s)-1 {
			res += strconv.Itoa(count) + string(c)
		}
	}
	return res
}
