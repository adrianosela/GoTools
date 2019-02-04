package sort

import (
	"github.com/adrianosela/GoTools/primitives/slices"
)

// MergeSort runs the merge sort algorithm on an array
// Running Time : Best/Avg/Worst Case = Θ(nlog(n))
func MergeSort(arr []int) {
	merge := func(s []int, lo, mid, hi int) {
		tmp := []int{}
		a, b := lo, mid+1
		for k := lo; k <= hi; k++ {
			if a <= mid && (b > hi || s[a] < s[b]) {
				tmp = append(tmp, s[a])
				a++
			} else {
				tmp = append(tmp, s[b])
				b++
			}
		}
		for k := lo; k <= hi; k++ {
			s[k] = tmp[k-lo]
		}
	}

	var msort func([]int, int, int)
	msort = func(a []int, lo, hi int) {
		if lo >= hi {
			return
		}
		mid := (lo + hi) / 2
		msort(a, lo, mid)
		msort(a, mid+1, hi)
		merge(a, lo, mid, hi)
	}

	msort(arr, 0, len(arr)-1) // sort between indices 0 and n-1
}

// InsertionSort runs the insertion sort algorithm on an array
// Running Time : Best Case = Θ(n), Avg/Worst Case = Θ(n^2)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		var j int
		val := arr[i]
		for j = i; j > 0 && arr[j-1] > val; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = val
	}
}

// SelectionSort runs the selection sort algorithm on an array
// Running Time : Best/Avg/Worst Case = Θ(n^2)
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// swap current index for smallest ahead in array
		idxOfMin, _ := slices.FindSmallest(arr, i)
		arr[i], arr[idxOfMin] = arr[idxOfMin], arr[i]
	}
}
