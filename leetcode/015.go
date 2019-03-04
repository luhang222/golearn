package leetcode

func ThreeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	var neg = make(map[int]int, 0)
	var pos = make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			neg[nums[i]]++
		} else {
			pos[nums[i]]++
		}
	}
	if n, e := pos[0]; e {
		if n >= 3 {
			res = append(res, []int{0, 0, 0})
		}
	}
	for k := range neg {
		skip := make(map[int]bool, 0)
		for kk := range pos {
			if _, e := skip[kk]; e {
				continue
			}
			dif := -k - kk
			if dif < 0 {
				n, e := neg[dif]
				if e && n > 0 && (dif != k || (dif == k && n > 1)) {
					res = append(res, []int{k, kk, dif})
				}
			} else {
				n, e := pos[dif]
				if e && (dif != kk || (dif == kk && n > 1)) {
					skip[dif] = true
					res = append(res, []int{k, kk, dif})
				}
			}
		}
		neg[k] = 0
	}
	return res
}
