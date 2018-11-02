package test

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := GenRandomSlice(1000)
	BubbleSort(arr)
	if !checkSort(arr) {
		t.Error("incorrect result!")
	}
}

func TestInsertSort(t *testing.T) {
	arr := GenRandomSlice(1000)
	InsertSort(arr)
	if !checkSort(arr) {
		t.Error("incorrect result!")
	}
}

func TestQuickSort(t *testing.T) {
	arr := GenRandomSlice(1000)
	QuickSort(arr)
	if !checkSort(arr) {
		t.Error("incorrect result!")
	}
}

func TestHeapSort(t *testing.T) {
	arr := GenRandomSlice(1000)
	HeapSort(arr)
	if !checkSort(arr) {
		t.Error("incorrect result!")
	}
}

func checkSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
