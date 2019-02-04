package slices

// FindSmallest returns the index and value of the smallest value starting at index b
func FindSmallest(arr []int, b int) (int, int) {
	r := b
	min := arr[b]
	for i := b + 1; i < len(arr); i++ {
		if arr[i] < min {
			r = i
			min = arr[i]
		}
	}
	return r, min
}
