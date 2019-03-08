package leetcode

func NextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	swapA, swapB, dif := -1, 0, 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			swapA = i
			dif = nums[i+1] - nums[i]
			break
		}
	}
	if swapA != -1 {
		for i := swapA + 1; i <= len(nums)-1; i++ {
			if nums[i] > nums[swapA] && nums[i]-nums[swapA] <= dif {
				dif = nums[i] - nums[swapA]
				swapB = i
			}
		}
		nums[swapA], nums[swapB] = nums[swapB], nums[swapA]
	}
	i, j := swapA+1, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return
}
