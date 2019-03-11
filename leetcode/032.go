package leetcode

func LongestValidParentheses(s string) int {
	llnum, lrnum, rlnum, rrnum := 0, 0, 0, 0
	max := 0
	length := len(s) - 1
	for i := 0; i <= length; i++ {
		if s[i] == 40 {
			llnum++
		}
		if s[i] == 41 {
			lrnum++
		}
		if s[length-i] == 40 {
			rlnum++
		}
		if s[length-i] == 41 {
			rrnum++
		}
		if lrnum > llnum {
			llnum, lrnum = 0, 0
		} else if lrnum == llnum {
			if lrnum > max {
				max = lrnum
			}
		}
		if rlnum > rrnum {
			rlnum, rrnum = 0, 0
		} else if rlnum == rrnum {
			if rlnum > max {
				max = rlnum
			}
		}
	}
	return max * 2
}
