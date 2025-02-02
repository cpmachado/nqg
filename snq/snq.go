package snq

import (
	"fmt"
	"io"
	"strings"
)

// Snq implements the original Simple N-Queen puzzle solver
func Snq(w io.Writer, pos []int, k int) {
	if k == len(pos) {
		printPosition(w, pos)
		return
	}

	n := len(pos)
	for i := 0; i < n; i++ {
		pos[k] = i
		if isValidPosition(pos, k) {
			Snq(w, pos, k+1)
		}
	}
}

func iabs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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

func printPosition(w io.Writer, pos []int) {
	var sb strings.Builder

	for _, p := range pos {
		fmt.Fprintf(&sb, "%v", p)
	}

	fmt.Fprintf(w, "%s\n", sb.String())
}
