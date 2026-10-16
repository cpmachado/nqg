package nqg

import (
	"io"
)

// SNQ implements the original Simple N-Queen puzzle solver
func SNQ(w io.Writer, pos []int, k int) {
	if k == len(pos) {
		printPosition(w, pos)
		return
	}

	n := len(pos)
	for i := range n {
		pos[k] = i
		if isValidPosition(pos, k) {
			SNQ(w, pos, k+1)
		}
	}
}

func isValidPosition(pos []int, k int) bool {
	p := pos[k]

	for i, v := range pos[:k] {
		d := iabs(p - v)
		if v == p || d == (k-i) {
			return false
		}
	}
	return true
}
