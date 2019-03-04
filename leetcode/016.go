package leetcode

import "math"

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	var lastDif int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if math.Abs(float64(sum-target)) < math.Abs(float64(lastDif)) || k == 2 {
					lastDif = sum - target
				}
			}
		}
	}
	return lastDif + target
}
