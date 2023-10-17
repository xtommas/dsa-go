package main

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(arr, lo, hi)

	qs(arr, lo, pivot-1)
	qs(arr, pivot+1, hi)
}

func quicksort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
