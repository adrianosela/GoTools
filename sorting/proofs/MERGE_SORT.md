### MergeSort

```
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
```

##### Proof By Induction:
// TODO

##### Runtime Complexity Analysis:

* **BEST/AVG/WORST CASE:** The problem is split in half by every recursive call to ```msort()```; then a call to ```merge()``` combines the two split (and now sorted) halves:
	* ```msort()``` has a runtime of Θ(lg(n)) as it halves the problem size each time. 
	* ```merge()``` has a runtime of Θ(n) as it iterates through a portion of the array.
	* The total runtime is Θ(nlg(n))

* **PROVING WITH RECURRENCE RELATION:** 
	* The base case is a constant time "return", i.e. **T(1) <= b**, where b is a positive constant
	* The standard case is splitting the problem in half and eventually reaching the base case, this yields the recurrence relation **T(n) <= 2T(n/2) + b*n**
	* Solving the recurrence relation: 
		* T(n) <= 2T(n/2) + bn
		* T(n) <= 2(2T(n/4)+b(n/2))+bn
		* T(n) <= 4T(n/4)+2bn
		* T(n) <= 4(2T(n/8)+b(n/8))+2bn
		* T(n) <= 8T(n/8)+3bn
		* **T(n) <= (2<sup>k</sup>)T(n/2<sup>k</sup>)+kbn**
		* Let n = 2<sup>k</sup> --> T(n) = nT(1)+lg(n)bn
		* So we have that the runtime of mergeSort is asymptotically bounded by Θ(nlog(n)), completing the proof	
