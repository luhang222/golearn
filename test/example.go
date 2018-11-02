package test

import (
	"math/rand"
	"time"
)

func BubbleSort(in []int) {
	for i := 1; i < len(in); i++ {
		for j := 0; j < len(in)-i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
}

func InsertSort(in []int) {
	for i := 0; i < len(in); i++ {
		min := i
		for j := i + 1; j < len(in); j++ {
			if in[i] > in[j] {
				min = j
			}
		}
		in[i], in[min] = in[min], in[i]
	}
}

func QuickSort(in []int) {
	if len(in) == 1 || len(in) == 0 {
		return
	}
	flag := 0
	v := in[flag]
	i := 0
	j := len(in) - 1
	for {
		for in[j] >= v && i < j {
			j--
		}
		for in[i] <= v && i < j {
			i++
		}
		if i >= j {
			break
		}
		in[i], in[j] = in[j], in[i]
	}
	in[flag], in[i] = in[i], in[flag]
	QuickSort(in[:i])
	QuickSort(in[i+1:])
}

func HeapSort(in []int) {
	if len(in) == 1 {
		return
	}
	for i := len(in) - 1; i > 0; i-- {
		if i%2 == 1 {
			if in[i] < in[(i-1)/2] {
				in[i], in[(i-1)/2] = in[(i-1)/2], in[i]
			}
		} else {
			if in[i] < in[(i/2-1)] {
				in[i], in[(i/2-1)] = in[(i/2-1)], in[i]
			}
		}
	}
	for i := 1; i < len(in)-1; i++ {
		rebuildHeap(in[i:])
	}
}

func rebuildHeap(in []int) {
	if len(in) == 1 {
		return
	}
	top := 0
	for top < len(in)-1 {
		if top*2+1 < len(in) && in[top] > in[top*2+1] {
			in[top], in[top*2+1] = in[top*2+1], in[top]
			top = top*2 + 1
		} else if top*2+2 < len(in) && in[top] > in[top*2+2] {
			in[top], in[top*2+2] = in[top*2+2], in[top]
			top = top*2 + 2
		} else {
			return
		}
	}
}

func GenRandomSlice(length int) []int {
	ret := make([]int, length)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < length; i++ {
		ret[i] = r.Intn(100000)
	}
	return ret
}
