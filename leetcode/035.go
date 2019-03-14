package leetcode

func SearchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if target <= nums[0] {
		return 0
	}
	i, j := 0, l
	for i < j {
		m := (i + j) / 2
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			i = m
		} else {
			j = m
		}
		if i+1 == j {
			return j
		}
	}
	return 0
}
