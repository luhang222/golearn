package leetcode

func RomanToInt(s string) int {
	res := 0
	xm := map[byte]int{
		77: 1000,
		68: 500,
		67: 100,
		76: 50,
		88: 10,
		86: 5,
		73: 1,
	}
	for i := len(s) - 1; i >= 0; i-- {
		if i > 0 && xm[s[i]] > xm[s[i-1]] {
			res += xm[s[i]] - xm[s[i-1]]
			i--
			continue
		}
		res += xm[s[i]]
	}
	return res
}
