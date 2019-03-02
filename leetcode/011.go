package leetcode

import "math"

func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var max int
	for i < j {
		a := float64(height[i])
		b := float64(height[j])
		if max < (j-i)*int(math.Min(a, b)) {
			max = (j - i) * int(math.Min(a, b))
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
