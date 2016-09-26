package matrix

// RowIgnore blacklists rows of a matrix from an operation. It's used by GeneratePartialIdentity to leave empty rows in
// a matrix.
type RowIgnore func(int) bool

// IgnoreNoRows implements the RowIgnore interface. It sets no rows to be blacklisted.
func IgnoreNoRows(row int) bool {
	return false
}

// IgnoreRows returns an impementation of the RowIgnore interface which is true at all given positions and false at all
// others.
func IgnoreRows(positions ...int) RowIgnore {
	return func(row int) bool {
		for _, cand := range positions {
			if row == cand {
				return true
			}
		}

		return false
	}
}
