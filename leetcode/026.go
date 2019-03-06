package leetcode

func removeDuplicates(nums []int) int {
	var max int
	var j int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[i]
			j++
			continue
		}
		if nums[i] > max {
			max = nums[i]
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
