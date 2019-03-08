package leetcode

import "strings"

/**
偷懒法
*/
func strStr_lazy(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

/**
sunday算法
*/
func StrStr_sunday(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	needleMap := map[int32]int{}
	for i, c := range needle {
		needleMap[c] = i
	}
	plen := len(needle)
	i := 0
	for i < len(haystack) {
		if i+plen > len(haystack) {
			return -1
		}
		if haystack[i] == needle[0] {
			f := true
			for k := 1; k < plen; k++ {
				if haystack[i+k] != needle[k] {
					if i+plen >= len(haystack) {
						return -1
					}
					index, e := needleMap[int32(haystack[i+plen])]
					if !e {
						i += plen + 1
						f = false
						break
					} else {
						i += plen - index
						f = false
						break
					}
				}
			}
			if f {
				return i
			}
		} else {
			i++
		}
	}
	return -1
}
