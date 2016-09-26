package matrix

// gaussJordan reduces the matrix according to the Gauss-Jordan Method.  Returns the augment matrix, the transformed
// matrix, and the set set of free variables.
func (e Matrix) gaussJordan() (aug, f Matrix, frees []int) {
	out, in := e.Size()

	aug = GenerateIdentity(in)
	for x := in; x < out; x++ {
		aug = append(aug, NewRow(in))
	}

	f = e.Dup() // Duplicate e away so we don't mutate it.

	row, col := 0, 0

	for row < out && col < in {
		// Find a non-zero element to move into the pivot position.
		i := f.FindPivot(row, col)
		if i == -1 { // Failed to find a pivot.
			frees = append(frees, col)
			col++

			continue
		}

		// Move it into position.
		f[row], f[i] = f[i], f[row]
		aug[row], aug[i] = aug[i], aug[row]

		// Make top position equal 1.
		s := NewNumber().Inv(f[row][col])
		f[row] = f[row].MulScalar(s)
		aug[row] = aug[row].MulScalar(s)

		// Cancel out the (row)th position for every row above and below it.
		for j, _ := range f {
			if j != row && f[j][col].Sign() != 0 {
				s.Neg(f[j][col])

				aug[j] = aug[j].Add(aug[row].MulScalar(s))
				f[j] = f[j].Add(f[row].MulScalar(s))
			}
		}

		row++
		col++
	}

	// Add the rest of the free variables for completion.
	for x := out; x < in; x++ {
		frees = append(frees, x)
	}

	return
}

// NullSpace returns a basis for the matrix's nullspace.
func (e Matrix) NullSpace() (basis []Row) {
	out, in := e.Size()
	if out == 0 {
		return []Row{}
	}

	_, f, frees := e.gaussJordan()

	for _, free := range frees {
		input := NewRow(in)
		input[free].SetOne()

		for _, row := range f {
			if row[free].Sign() != 0 {
				input[row.Height()].Neg(row[free])
			}
		}

		basis = append(basis, input)
	}

	return
}
