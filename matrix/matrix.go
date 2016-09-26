// Package matrix implements basic operations on matrices in a field and the generation of new ones.
package matrix

// Matrix is a matrix.
type Matrix []Row

// Mul right-multiplies a matrix by a row.
func (e Matrix) Mul(f Row) Row {
	out, in := e.Size()
	if in != len(f) {
		panic("Can't multiply by row that is wrong size!")
	}

	res := NewRow(out)
	for i, row := range e {
		res[i].Set(row.DotProduct(f))
	}

	return res
}

// Add adds two matrices.
func (e Matrix) Add(f Matrix) Matrix {
	a, _ := e.Size()

	out := make([]Row, a)
	for i, _ := range out {
		out[i] = e[i].Add(f[i])
	}

	return out
}

// Compose returns the result of composing e with f.
func (e Matrix) Compose(f Matrix) Matrix {
	n, m := e.Size()
	p, q := f.Size()

	if m != p {
		panic("Can't multiply matrices of wrong size!")
	}

	out := GenerateEmpty(n, q)
	g := f.Transpose()

	for i, e_i := range e {
		for j, g_j := range g {
			out[i][j].Set(e_i.DotProduct(g_j))
		}
	}

	return out
}

// Invert computes the multiplicative inverse of a matrix, if it exists.
func (e Matrix) Invert() (Matrix, bool) {
	inv, _, frees := e.gaussJordan()
	return inv, len(frees) == 0
}

// Transpose returns the transpose of a matrix.
func (e Matrix) Transpose() Matrix {
	n, m := e.Size()
	out := GenerateEmpty(m, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			out[j][i].Set(e[i][j])
		}
	}

	return out
}

// Trace returns the trace of a matrix.
func (e Matrix) Trace() (out Number) {
	for i, e_i := range e {
		out.Add(out, e_i[i])
	}

	return
}

// FindPivot finds a row with non-zero entry in column col, starting at the given row and moving down. It returns the
// index of the given row or -1 if one does not exist.
func (e Matrix) FindPivot(row, col int) int {
	for i, e_i := range e[row:] {
		if e_i[col].Sign() != 0 {
			return row + i
		}
	}

	return -1
}

// Equals returns true if two matrices are equal and false otherwise.
func (e Matrix) Equals(f Matrix) bool {
	a, _ := e.Size()
	c, _ := f.Size()

	if a != c {
		return false
	}

	for i, _ := range e {
		if !e[i].Equals(f[i]) {
			return false
		}
	}

	return true
}

// Dup returns a duplicate of this matrix.
func (e Matrix) Dup() Matrix {
	n, m := e.Size()
	f := GenerateEmpty(n, m)

	for i, _ := range e {
		for j, _ := range e[i] {
			f[i][j].Set(e[i][j])
		}
	}

	return f
}

// Size returns the dimensions of the matrix in (Rows, Columns) order.
func (e Matrix) Size() (int, int) {
	if len(e) == 0 {
		return 0, 0
	} else {
		return len(e), len(e[0])
	}
}
