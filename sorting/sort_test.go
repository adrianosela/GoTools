package sort

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	Convey("Testing Sorting Functions", t, func() {
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
				Convey(ts.name, func() {
					So(tmp, ShouldResemble, ts.ordered)
				})
			}
		})
		Convey("Test InsertionSort", func() {
			for _, ts := range tests {
				tmp := make([]int, len(ts.disordered))
				copy(tmp, ts.disordered)
				InsertionSort(tmp)
				Convey(ts.name, func() {
					So(tmp, ShouldResemble, ts.ordered)
				})
			}
		})
		Convey("Test SelectionSort", func() {
			for _, ts := range tests {
				tmp := make([]int, len(ts.disordered))
				copy(tmp, ts.disordered)
				SelectionSort(tmp)
				Convey(ts.name, func() {
					So(tmp, ShouldResemble, ts.ordered)
				})
			}
		})
		// Convey("RuntimeAnalysys", func() {
		// 	largeArr := make([]int, 50)
		// 	runAlgorithmRuntimeTest(largeArr)
		// })
	})
}

func runAlgorithmRuntimeTest(arr []int) {
	ms, ss, is := make([]int, len(arr)), make([]int, len(arr)), make([]int, len(arr))
	copy(ms, arr)
	copy(ss, arr)
	copy(is, arr)

	msDone := make(chan bool)
	ssDone := make(chan bool)
	isDone := make(chan bool)

	go func() {
		MergeSort(ms)
		msDone <- true
	}()
	go func() {
		SelectionSort(ss)
		ssDone <- true
	}()
	go func() {
		InsertionSort(is)
		isDone <- true
	}()

	done := 0
	for done != 3 {
		select {
		case _ = <-msDone:
			done++
			fmt.Println("MergeSort done")
		case _ = <-ssDone:
			done++
			fmt.Println("InsertionSort done")
		case _ = <-isDone:
			done++
			fmt.Println("SelectionSort done")
		}
	}

}
