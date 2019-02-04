### SelectionSort

```
// SelectionSort runs the selection sort algorithm on an array
// Running Time : Best/Avg/Worst Case = Θ(n^2)
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// swap current index for smallest ahead in array
		idxOfMin, _ := slices.FindSmallest(arr, i)
		arr[i], arr[idxOfMin] = arr[idxOfMin], arr[i]
	}
}
```

##### Proof By Induction:

* **Theorem:** At the beginning of iteration i, arr[0..i-1] contains the smallest i elements in the array sorted in ascending order.

* **Proof:**
	* **BASE CASE:** At the beginning of iteration i = 0 (before any code has executed), arr[0..i-1] has no elements, and thus the theorem holds for the base case
	* **INDUCTIVE HYPOTHESIS:** Assume the theorem holds for some arbitrary iteration i
	* **INDUCTIVE STEP:** Then at the beginning of iteration i, we find the minimum of arr[i..n-1] and swap it for arr[i]. Then arr[i] >= all elements in arr[0..i-1] - by the inductive hypothesis - and arr[i] <= arr[i..n-1] by the correctness of the ```FindSmallest(A []int, b int)``` function.
	* **TERMINATION:** Hence at the end of iteration i+1, the elements in arr[0..(i+1)-1] = arr[0..i] are sorted in ascending order and the theorem holds for.

##### Runtime Complexity Analysis:

* **BEST/AVG/WORST CASE:** In all cases, even if the input is sorted, we iterate through the array twice, once as we traverse the array, and once to find the min from arr[i..n-1] in each iteration of the outer for-loop. This yields a runtime of Θ(n^2)