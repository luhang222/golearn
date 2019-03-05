package leetcode

import "sort"

func FourSum(nums []int, target int) [][]int {
	var fourSum [][]int
	sort.Ints(nums)
	var len = len(nums)
	if len <= 3 {
		return nil
	}

	for i := 0; i < len; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var target1 = target - nums[i]
		for j := i + 1; j < len; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			var target2 = target1 - nums[j]
			l := j + 1
			r := len - 1
			for l < r {
				var tempSum = nums[l] + nums[r]
				if tempSum == target2 {
					fourSum = append(fourSum, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if tempSum < target2 {
					l++
				} else {
					r--
				}
			}
		}
	}
	return fourSum
}
