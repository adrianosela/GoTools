### InsertionSort

```
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
```

##### Proof By Induction:

* **Theorem:** At the beginning of iteration i of the outer for-loop, arr[0..i-1] is sorted in ascending order

* **Proof:**
	* **BASE CASE:** At the beginning of iteration i=0 of the for-loop (before any code executes), arr[0..i-1] contains no elements, and thus the theorem holds true.
	* **INDUCTIVE HYPOTHESIS:** We assume that the invariant is true at the beginning of iteration i
	* **INDUCTIVE STEP:** At iteration i, the inner for-loop will iteratively swap arr[j=i] with arr[largest] where largest is the index of the largest value in arr[0..j]. The inner for-loop executes until arr[i] holds the largest value in arr[0..i]. This leaves the elements in arr[0..i-1] sorted in ascending order and thus the theorem holds.
	* **TERMINATION:** At the end of iteration i = n-1, (or what would be the start of iteration n) all of arr[0..n-1] are sorted in ascending order.

##### Runtime Complexity Analysis:

* **BEST CASE:** If the input array is already sorted, the inner loop will not execute at all, and we simply iterate once through the entire array, yielding Θ(n)


* **AVG/WORST CASE:** On average, and at the worst case we iterate through the loop twice, yielding Θ(n^2)