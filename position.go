package nqg

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
