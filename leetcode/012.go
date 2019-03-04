package leetcode

func IntToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	res := ""
	x := 1000
	xm := map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}
	for x > 0 {
		s, ss := num/x, num%x
		if s <= 3 {
			for i := 0; i < s; i++ {
				res += xm[x]
			}
		} else if s == 4 {
			res += xm[x] + xm[x*5]
		} else if s <= 8 {
			res += xm[x*5]
			for i := 0; i < s-5; i++ {
				res += xm[x]
			}
		} else {
			res += xm[x] + xm[x*10]
		}
		num = ss
		x /= 10
	}
	return res
}
