package main

import (
	"fmt"
	"testing"
)

func Test_heap_sort(t *testing.T) {
	array := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	HeapSort(array)
	fmt.Println(array)
}
