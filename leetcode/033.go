package leetcode

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	n := findN(nums)
	res := -1
	if n == 0 || nums[0] > target {
		res = find(nums[n:], target)
		if res != -1 {
			res = res + n
		}
	} else {
		res = find(nums[0:n], target)
	}
	return res
}

func find(nums []int, target int) int {
	l := len(nums)
	if l == 1 {
		if nums[0] != target {
			return -1
		} else {
			return 0
		}
	}
	index := 0
	if l&1 == 0 {
		index = l / 2
	} else {
		index = (l - 1) / 2
	}
	if nums[index] == target {
		return index
	}
	if nums[index] > target {
		return find(nums[:index], target)
	} else {
		tmp := find(nums[index:], target)
		if tmp != -1 {
			return tmp + index
		} else {
			return tmp
		}
	}
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
