package disjointsets

// DisjointSets is a functional representation of disjoint sets
type DisjointSets interface {
	// joins the sets containing values k1 and k2
	Union(k1, k2 interface{})
	// returns the representative element of the set containing item k
	Find(k interface{}) interface{}
}

// ADJSets is an array representation of disjoint sets
//
// Example representation of S = {{0,1,4},{2,7},{3,5,6}}
// ADJSets {
//  elems: []int{0,1,2,3,4,5,6,7},
//  reps:  []int{0,0,2,3,0,3,3,2},
// }
//
// Runtime Analysis of Implementation:
// Find - O(1)
// Union - Θ(n) (have to check all values in reps array)
type ADJSets struct {
	elems []int
	reps  []int
}

// Find returns the representative element of the set containing k
func (s *ADJSets) Find(k int) int {
	return s.reps[k]
}

// Union points the representative element of one set to that of another
// thus joining the two sets, and having the second element be the
// representative element of the new set containing all elements in both sets
func (s *ADJSets) Union(k1, k2 int) {
	s.reps[s.Find(k1)] = s.Find(k2)
}
