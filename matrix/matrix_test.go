package matrix

import (
	"testing"
)

var (
	testM = Matrix{
		Row{NewRatI(1), NewRatI(2), NewRatI(0)},
		Row{NewRatI(0), NewRatI(3), NewRatI(1)},
		Row{NewRatI(4), NewRatI(0), NewRatI(1)},
	}
	testSingular = Matrix{
		Row{NewRatI(1), NewRatI(2), NewRatI(1)},
		Row{NewRatI(0), NewRatI(3), NewRatI(1)},
		Row{NewRatI(1), NewRatI(5), NewRatI(2)},
	}
)

// TestMatrixMul checks correctness of right-multiplication.
func TestMatrixMul(t *testing.T) {
	in := Row{NewRatI(7), NewRatI(9), NewRatI(8)}
	out := Row{NewRatI(7 + 2*9), NewRatI(3*9 + 8), NewRatI(4*7 + 8)}

	cand := testM.Mul(in)
	if !out.Equals(cand) {
		t.Fatal("candidate not equal to correct output")
	}
}

// TestInverse checks that calling Invert on an invertible matrix returns true, and that composing this matrix with the
// original returns the identity.
func TestInverse(t *testing.T) {
	inv, ok := testM.Invert()
	if !ok {
		t.Fatal("failed to invert invertible matrix")
	}

	id := GenerateIdentity(3)
	cand1, cand2 := testM.Compose(inv), inv.Compose(testM)
	if !cand1.Equals(id) || !cand2.Equals(id) {
		t.Fatal("matrix times inverse is not identity")
	}
}

// TestSingular checks that calling Invert on a singular matrix returns false.
func TestSingular(t *testing.T) {
	_, ok := testSingular.Invert()
	if ok {
		t.Fatal("inverted singular matrix")
	}

	basis := testSingular.NullSpace()
	if len(basis) != 1 {
		t.Fatal("basis is wrong size")
	} else if basis[0].IsZero() {
		t.Fatal("vector is zero")
	}

	zero := testSingular.Mul(basis[0])
	if !zero.IsZero() {
		t.Fatal("not in nullspace")
	}
}
