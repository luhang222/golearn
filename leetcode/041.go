package leetcode

import "fmt"

func FirstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		} else {
			i++
		}
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}
