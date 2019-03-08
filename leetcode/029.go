package leetcode

func Divide_lazy(dividend int, divisor int) int {
	i := dividend / divisor
	if i > 2147483647 || i < -2147483648 {
		return 2147483647
	}
	return i
}

//通过右移确定 n * 2 << 0 , m * 2 << 1, o * 2 << 2.....
