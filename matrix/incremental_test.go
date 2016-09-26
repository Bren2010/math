package matrix

import (
	"testing"
)

// TestIncrementalMatrix checks that adding linearly independent rows to a matrix succeeds and recovers the original
// matrix.
func TestIncrementalMatrix(t *testing.T) {
	im := NewIncrementalMatrix(3)
	ok1 := im.Add(testM[0])
	ok2 := im.Add(testM[1])
	ok3 := im.Add(testM[2])

	if !(ok1 && ok2 && ok3) {
		t.Fatal("rows are linearly dependent")
	} else if !im.FullyDefined() {
		t.Fatal("matrix is not fully defined")
	} else if !im.Matrix().Equals(testM) {
		t.Fatal("not equal to original")
	}
}

// TestIncrementalMatrixSingular checks that adding linearly dependent rows to a matrix results in failure.
func TestIncrementalMatrixSingular(t *testing.T) {
	im := NewIncrementalMatrix(3)
	ok1 := im.Add(testSingular[0])
	ok2 := im.Add(testSingular[1])
	ok3 := im.Add(testSingular[2])

	if !(ok1 && ok2) {
		t.Fatal("rows are linearly dependent")
	} else if ok3 || !im.IsIn(testSingular[2]) {
		t.Fatal("row is linearly independent")
	} else if im.FullyDefined() {
		t.Fatal("matrix is fully defined")
	}
}

// TestIncrementalNovel checks that NovelRow returns a linearly independent vector.
func TestIncrementalNovel(t *testing.T) {
	im := NewIncrementalMatrix(3)
	im.Add(testSingular[0])
	im.Add(testSingular[1])

	r := im.NovelRow()
	if im.IsIn(r) {
		t.Fatal("row is already in matrix")
	}
}
