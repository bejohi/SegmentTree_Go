package segTree

import "math"

// ISegmentTree defines a simple segment tree with common
// functions.
type ISegmentTree interface {
	// Returns the result of the segment in the given range.
	// The type of the result (e.g. sum, min/ max) depends on the implementation.
	Query(startPosition int, endPosition int) float64

	// Updates the position and all related sub segments with the given newValue.
	Update(position int, newValue float64)

	// Constructs a segment tree from a given input array. The size of the
	// segment tree will be (2 * len(initArr)) + 3
	Construct(initArr *[]float64) error

	// Returns the complete tree representation.
	GetTree() (*[]float64, error)
}

// CalculateTreeSize returns the max. size of a segment tree for the given input array length
func CalculateTreeSize(inputArrLength int) int {
	x := math.Ceil(math.Log(float64(inputArrLength)) / math.Log(2))
	return 2 * int(math.Pow(2, x)-1)
}
