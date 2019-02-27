package leetcode

import (
	"math"
)

func Reverse(x int) int {
	f := x < 0
	x = int(math.Abs(float64(x)))
	var res int
	bit := math.Floor(math.Log10(float64(x))) + 1
	var i float64 = 1
	for {
		single := x % 10
		x = x / 10
		res += int(math.Pow(10, bit-i) * float64(single))
		if (f && float64(res) > math.Pow(2, 31)) || (!f && float64(res) > (math.Pow(2, 31))-1) {
			return 0
		}
		i++
		if x == 0 {
			break
		}
	}
	if f {
		res -= 2 * res
	}
	return res
}
