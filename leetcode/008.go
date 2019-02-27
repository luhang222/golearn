package leetcode

func MyAtoi(str string) int {
	b := []rune(str)
	var last rune
	var res int
	var f = false
	for _, bt := range b {
		if bt == 32 {
			if last == 0 {
				last = bt
				continue
			}
			if last != 32 {
				break
			}
		} else if bt >= 48 && bt <= 57 {
			if last == 0 || last == 32 {
				last = bt
				res = int(bt) - 48
				continue
			}
			res = res*10 + int(bt) - 48
			if f && res > 2147483648 {
				return -2147483648
			}
			if !f && res > 2147483647 {
				return 2147483647
			}
		} else if bt == 45 {
			if last == 0 || last == 32 {
				f = true
				last = bt
				continue
			}
			break
		} else if bt == 43 {
			if last == 0 || last == 32 {
				f = false
				last = bt
				continue
			}
			break
		} else {
			if last == 0 || last == 32 {
				return 0
			}
			break
		}
	}
	if f {
		res -= 2 * res
	}
	return res
}
