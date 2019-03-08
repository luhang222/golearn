package leetcode

import (
	"strings"
)

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 || len(s) < len(words[0])*len(words) {
		return res
	}
	wl := len(words[0])
	indexMap := make(map[int]string, 0)
	countMap := make(map[string]int, 0)
	for _, w := range words {
		countMap[w]++
		i := strings.Index(s, w)
		if i == -1 {
			return res
		} else {
			indexMap[i] = w
			for {
				j := strings.Index(s[i+1:], w)
				if j == -1 {
					break
				}
				i++
				indexMap[i+j] = w
			}
		}
	}
	for k, w := range indexMap {
		tmp := make(map[string]int, 0)
		for tk, tw := range countMap {
			tmp[tk] = tw
		}
		tmp[w]--
		if tmp[w] == 0 {
			delete(tmp, w)
		}
		wc := 1
		for wc < len(words) {
			kk := k + wl*wc
			if ww, e := indexMap[kk]; e {
				if _, ee := tmp[ww]; !ee {
					break
				}
				tmp[ww]--
				if tmp[ww] < 0 {
					break
				}
				if tmp[ww] == 0 {
					delete(tmp, ww)
				}
			} else {
				break
			}
			wc++
		}
		if len(tmp) == 0 {
			res = append(res, k)
		}
	}
	return res
}

/**
best
*/
func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(words) == 0 {
		return result
	}
	wordMap := make(map[string]int)
	for _, v := range words {
		wordMap[v] = wordMap[v] + 1
	}
	wordLen := len(words[0])
	window := len(words) * wordLen
	for i := 0; i < wordLen; i++ {
		for j := i; j+window <= len(s); j = j + wordLen {
			tmpStr := s[j : j+window]
			temMap := make(map[string]int)
			for k := len(words) - 1; k >= 0; k-- {
				//get the word from subStr
				word := tmpStr[k*wordLen : (k+1)*wordLen]
				count := temMap[word] + 1
				if count > wordMap[word] {
					+j = j + k*wordLen
					break
				} else if k == 0 {
					result = append(result, j)
				} else {
					temMap[word] = count
				}
			}
		}
	}
	return result
}
