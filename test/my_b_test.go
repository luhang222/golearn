package test

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BubbleSort(GenRandomSlice(10000))
	}
}

func BenchmarkInsertSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InsertSort(GenRandomSlice(10000))
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		QuickSort(GenRandomSlice(10000))
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		HeapSort(GenRandomSlice(10000))
	}
}
