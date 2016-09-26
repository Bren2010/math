package matrix

// GenerateIdentity generates the n-by-n identity matrix.
func GenerateIdentity(n int) Matrix {
	return GeneratePartialIdentity(n, IgnoreNoRows)
}

// GeneratePartialIdentity generates the n-by-n identity matrix on some rows and leaves others zero (the rows where
// ignore(row) == true).
func GeneratePartialIdentity(n int, ignore RowIgnore) Matrix {
	out := GenerateEmpty(n, n)

	for i := 0; i < n; i++ {
		if !ignore(i) {
			out[i][i].SetOne()
		}
	}

	return out
}

// GenerateFull generates the n-by-n matrix with all entries set to 1.
func GenerateFull(n, m int) Matrix {
	out := GenerateEmpty(n, m)

	for i, _ := range out {
		for j, _ := range out[i] {
			out[i][j].SetOne()
		}
	}

	return out
}

// GenerateEmpty generates the n-by-n matrix with all entries set to 0.
func GenerateEmpty(n, m int) Matrix {
	out := make([]Row, n)

	for i := 0; i < n; i++ {
		out[i] = NewRow(m)
	}

	return Matrix(out)
}

// GeneratePermutationMatrix generates an n-by-n permutation matrix corresponding to a permutation of {0, ..., n-1}.
func GeneratePermutationMatrix(permutation []int) Matrix {
	n := len(permutation)
	out := GenerateEmpty(n, n)

	for i, j := range permutation {
		out[j][i].SetOne()
	}

	return out
}
