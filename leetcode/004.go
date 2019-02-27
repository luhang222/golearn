package leetcode

/**
时间复杂度为O(m+n)，题目要求O(log(m+n)) ...
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	merged := make([]int, 0, len(nums1)+len(nums2))
	if len(nums1) == 0 {
		merged = nums2
	} else if len(nums2) == 0 {
		merged = nums1
	} else {
		for {
			if i == len(nums1) {
				merged = append(merged, nums2[j:]...)
				break
			}
			if j == len(nums2) {
				merged = append(merged, nums1[i:]...)
				break
			}
			if nums1[i] < nums2[j] {
				merged = append(merged, nums1[i])
				i++
			} else {
				merged = append(merged, nums2[j])
				j++
			}
		}
	}
	if len(merged)%2 == 1 {
		return float64(merged[(len(merged)-1)/2])
	} else {
		return float64(merged[len(merged)/2]+merged[len(merged)/2-1]) / 2
	}
}
