package leetcode

func Search(nums []int, target int) int {
	n := findN(nums)
	return n
}

func findN(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	index := 0
	if len(nums)&1 == 0 {
		index = l / 2
	} else {
		index = (l - 1) / 2
	}
	if nums[index] < nums[index-1] {
		return index
	} else {
		a := findN(nums[:index])
		if a != 0 {
			return a
		} else {
			b := findN(nums[index:])
			if b != 0 {
				return b + index
			} else {
				return b
			}
		}
	}
}
