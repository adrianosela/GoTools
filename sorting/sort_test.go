package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	Convey("Testing Sorting Functions", t, func() {
		// Tests will be ordering the following arrays:
		tests := []struct {
			name       string
			disordered []int
			ordered    []int
		}{
			{
				name:       "Empty Slice",
				disordered: []int{},
				ordered:    []int{},
			},
			{
				name:       "Odd Slice Length",
				disordered: []int{7, 2, 15, 99, 12},
				ordered:    []int{2, 7, 12, 15, 99},
			},
			{
				name:       "Even Slice Length",
				disordered: []int{3, 7, 14, 342, 74, 9, 1, 234, 563, 2, 15, 99, 12},
				ordered:    []int{1, 2, 3, 7, 9, 12, 14, 15, 74, 99, 234, 342, 563},
			},
			{
				name:       "Slice With Duplicates",
				disordered: []int{3, 7, 14, 342, 74, 9, 1, 14, 3, 234, 563, 2, 15, 99, 12},
				ordered:    []int{1, 2, 3, 3, 7, 9, 12, 14, 14, 15, 74, 99, 234, 342, 563},
			},
			{
				name:       "Slice With Negatives",
				disordered: []int{7, 2, 15, 99, -12},
				ordered:    []int{-12, 2, 7, 15, 99},
			},
		}
		Convey("Test MergeSort", func() {
			for _, ts := range tests {
				tmp := make([]int, len(ts.disordered))
				copy(tmp, ts.disordered)
				MergeSort(tmp)
				So(tmp, ShouldResemble, ts.ordered)
			}
		})
		Convey("Test InsertionSort", func() {
			for _, ts := range tests {
				tmp := make([]int, len(ts.disordered))
				copy(tmp, ts.disordered)
				InsertionSort(tmp)
				So(tmp, ShouldResemble, ts.ordered)
			}
		})
		Convey("Test SelectionSort", func() {
			for _, ts := range tests {
				tmp := make([]int, len(ts.disordered))
				copy(tmp, ts.disordered)
				SelectionSort(tmp)
				So(tmp, ShouldResemble, ts.ordered)
			}
		})
	})
}
