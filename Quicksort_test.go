package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{9, 3, 7, 4, 69, 420, 42}, []int{3, 4, 7, 9, 42, 69, 420}},
		{[]int{3, 7, 9, 1}, []int{1, 3, 7, 9}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.nums)
		t.Run(testname, func(t *testing.T) {
			quicksort(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("got %d, want %d", tt.nums, tt.want)
			}
		})
	}
}
