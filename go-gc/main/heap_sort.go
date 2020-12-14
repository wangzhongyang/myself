package main

import (
	"fmt"
)

func HeapSort(array []int) {
	m := len(array)
	s := m / 2
	for i := s; i > -1; i-- {
		heap(array, i, m-1)
	}
	fmt.Println(array)
	for i := m - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}
}

func heap(array []int, i, end int) {
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && array[r] > array[l] {
		n = r
	}
	if array[i] > array[n] {
		return
	}
	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}
