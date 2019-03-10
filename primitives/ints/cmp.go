package ints

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return x *-1
	} else {
		return x
	}
}

