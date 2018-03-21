package segTree

import (
	"errors"
)

// SegmentTreeSum is an implementation for the ISegmentTree interface,
// where the query operation will return the sum of the segment
type SegmentTreeSum struct {
	size          int
	offset        int
	isConstructed bool
	tree          []float64
}

func (segTree *SegmentTreeSum) Query(startPosition int, endPosition int) float64 {
	return 0
}

func (segTree *SegmentTreeSum) Update(position int, newValue float64) {

}

func (segTree *SegmentTreeSum) GetTree() (*[]float64, error) {
	if segTree.tree == nil {
		return nil, errors.New("TREE WAS NOT CONSTRUCTED")
	}

	newTree := make([]float64, len(segTree.tree))
	copy(newTree, segTree.tree)
	return &newTree, nil
}

func (segTree *SegmentTreeSum) recConstruct(initArr *[]float64, left int, right int, index int) float64 {
	if left == right {
		(segTree.tree)[index] = (*initArr)[left]
		return (*initArr)[left]
	}

	mid := int(left + (right-left)/2)
	(segTree.tree)[index] = segTree.recConstruct(initArr, left, mid, index*2+1) +
		segTree.recConstruct(initArr, mid+1, right, index*2+2)
	return (segTree.tree)[index]

}

func (segTree *SegmentTreeSum) Construct(initArr *[]float64) error {
	if initArr == nil || len(*initArr) == 0 {

		return errors.New("INPUT ARRAY WAS NOT WELL FORMED")
	}

	if segTree.isConstructed {
		return errors.New("SEG TREE WAS ALREADY CONSTRUCTED")
	}
	segTree.isConstructed = true
	arrLen := len(*initArr)
	segTree.tree = make([]float64, CalculateTreeSize(arrLen))
	segTree.recConstruct(initArr, 0, arrLen-1, 0)

	segTree.offset = 2
	for segTree.offset < arrLen+2 {
		segTree.offset *= 2
	}

	return nil
}
