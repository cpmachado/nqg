package nqg

import (
	"io"
)

// SNQ implements the original Simple N-Queen puzzle solver
func SNQ(w io.Writer, pos []int, k int) {
	snq(w, pos, k)
}

func snq(w io.Writer, pos []int, k int) {
	n := len(pos)
	if k == n {
		printPosition(w, pos)
		return
	}
	for i := range n {
		pos[k] = i
		if isValidPosition(pos, k) {
			snq(w, pos, k+1)
		}
	}
}
