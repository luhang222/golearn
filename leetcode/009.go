package leetcode

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	n := x
	res := 0
	for n > 0 {
		i := n % 10
		n /= 10
		res = res*10 + i
	}
	return res == x
}
