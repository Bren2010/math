package matrix

// Row is a row of a matrix / vector.
type Row []Number

// NewRow returns an empty n-component row.
func NewRow(n int) Row {
	out := Row(make([]Number, n))
	for i := 0; i < n; i++ {
		out[i] = NewNumber()
	}
	return out
}

// Add adds two vectors.
func (e Row) Add(f Row) Row {
	le, lf := len(e), len(f)
	if le != lf {
		panic("can't add rows that are different sizes")
	}

	out := NewRow(le)
	for i := 0; i < le; i++ {
		out[i].Add(e[i], f[i])
	}

	return Row(out)
}

// Mul component-wise multiplies two vectors.
func (e Row) Mul(f Row) Row {
	le, lf := len(e), len(f)
	if le != lf {
		panic("can't multiply rows that are different sizes")
	}

	out := NewRow(le)
	for i := 0; i < le; i++ {
		out[i].Mul(e[i], f[i])
	}

	return Row(out)
}

// MulScalar multiplies each component of the row by the same scalar.
func (e Row) MulScalar(f Number) Row {
	le := len(e)
	out := NewRow(le)
	for i := 0; i < le; i++ {
		out[i].Mul(e[i], f)
	}

	return out
}

// Weight returns the hamming weight of this row.
func (e Row) Weight() Number {
	out := NewNumber()
	for i := 0; i < len(e); i++ {
		out.Add(out, e[i])
	}

	return out
}

// DotProduct computes the dot product of two vectors.
func (e Row) DotProduct(f Row) Number {
	return e.Mul(f).Weight()
}

// Returns true if e should be used to cancel out a column in f.
func (e Row) Cancels(f Row) bool {
	for i, _ := range e {
		if e[i].Sign() != 0 {
			return f[i].Sign() != 0
		}
	}

	return false
}

// IsZero returns true if the row is identically zero.
func (e Row) IsZero() bool {
	for _, e_i := range e {
		if e_i.Sign() != 0 {
			return false
		}
	}

	return true
}

// Height returns the position of the first non-zero entry in the row, or -1 if the row is zero.
func (e Row) Height() int {
	for i := 0; i < len(e); i++ {
		if e[i].Sign() != 0 {
			return i
		}
	}

	return -1
}

// Equals returns true if two rows are equal and false otherwise.
func (e Row) Equals(f Row) bool {
	le, lf := len(e), len(f)
	if le != lf {
		return false
	}

	for i := 0; i < le; i++ {
		if e[i].Cmp(f[i]) != 0 {
			return false
		}
	}

	return true
}

// Cmp returns -1 if row g is "less than" row f, 0 if equal, and 1 if "greater than". If you use sort a permutation
// matrix according to this, you'll always get the identity matrix.
func (e Row) Cmp(f Row) int {
	le, lf := len(e), len(f)
	if le != lf {
		panic("can't compare rows that are different sizes")
	}

	for i := 0; i < le; i++ {
		cmp := e[i].Cmp(f[i])
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

// Dup returns a duplicate of this row.
func (e Row) Dup() Row {
	le := len(e)
	out := NewRow(le)

	for i := 0; i < le; i++ {
		out[i].Set(e[i])
	}

	return out
}
