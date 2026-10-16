package nqg

import (
	"fmt"
	"io"
	"strings"
)

func iabs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func printPosition(w io.Writer, pos []int) {
	var sb strings.Builder

	for _, p := range pos {
		fmt.Fprintf(&sb, "%v", p)
	}

	fmt.Fprintf(w, "%s\n", sb.String())
}
