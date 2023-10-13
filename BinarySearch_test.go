package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{2, 3, 5, 7, 9, 43, 56, 67}, 43, 5},
		{[]int{2, 3, 5, 7, 9, 43, 56, 67}, 7, 3},
		{[]int{2, 3, 5, 7, 9, 43, 56, 67}, 26, -1},
		{[]int{2, 3}, 26, -1},
		{[]int{2, 3}, 2, 0},
		{[]int{2}, 2, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := binarySearch(tt.nums, tt.target)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
